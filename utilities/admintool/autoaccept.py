import requests
import time
import logging
import os

SLEEP_INTERVAL = 180 # Time between accept cycles, in seconds
RATE_LIMIT = 0.10 # Time to wait between rate-limit-sensitive API requests, in seconds

class AutoAccept():
    def __init__(self, admin_jwt: str):
        self.headers = {
            "Content-Type": "application/json",
            "Authorization": admin_jwt,
        }
        
        self.setup_logger()
        self.logger.info("Initializing script.")

        self.RATE_LIMIT = RATE_LIMIT  
        self.base_url = "https://api.hackillinois.org"

    def setup_logger(self) -> None:
        '''
            Set up logging to console and the log file.
        '''
        self.logger = logging.Logger("AutoAccept")

        # Console handler
        ch = logging.StreamHandler()
        ch.setLevel(logging.DEBUG)

        # File handler
        fh = logging.FileHandler("./AutoAccept.log")
        fh.setLevel(logging.DEBUG)

        formatter = logging.Formatter('%(asctime)s - %(levelname)s - %(message)s')
        ch.setFormatter(formatter)
        fh.setFormatter(formatter)

        self.logger.addHandler(ch)
        self.logger.addHandler(fh)

    def get_registered_users(self) -> set:
        '''
            Returns a set of all the users who have registered to the event.
            - This will also contain the users that have been accepted already.
        '''
        r = requests.get(f"{self.base_url}/mail/list/registered_users/", headers=self.headers)
        r.raise_for_status()
        return set(r.json()["userIds"])

    def get_accepted_users(self) -> set:
        '''
            Returns a set of all the users who have already been accepted to the event.
        '''
        r = requests.get(f"{self.base_url}/mail/list/accepted_wave_1/", headers=self.headers)
        r.raise_for_status()
        return set(r.json()["userIds"])

    def accept_user(self, id: str) -> bool:
        '''
            Accepts a user to the event with `id`. Returns a bool representing if the accept was successful. 
        '''
        json = {
            "id": id,
            "status": "ACCEPTED",
            "wave": 1
        }
        r = requests.post(f"{self.base_url}/decision/", headers=self.headers, json=json)
        return r.ok

    def finalize_user(self, id: str) -> bool:
        '''
            Finalizes a user's decision. Returns a bool representing if the finalization was successful.
        '''
        json = {
            "id": id,
            "finalized": True
        }

        r = requests.post(f"{self.base_url}/decision/finalize/", headers=self.headers, json=json)
        return r.ok

    def send_email(self, ids: list, template_id: str) -> bool:
        '''
            Sends acceptance email to users in the `ids` list.
        '''
        json = {
            "ids": ids,
            "template": template_id
        }

        r = requests.post(f"{self.base_url}/mail/send/", headers=self.headers, json=json)
        return r.ok

    def run_accept_cycle(self) -> None:
        # 1. Get a list of all registrations and already accepted users
        registered_users = self.get_registered_users()
        accepted_users = self.get_accepted_users()

        # 2. Subtract the list of users that have already been accepted (accepted_wave_1 Mail List)
        users_to_accept = registered_users - accepted_users
        users_accepted = set()
        users_finalized = set()

        # 3. Accept all users in this list (RATE LIMITED)
        for user in users_to_accept:
            if self.accept_user(user):
                users_accepted.add(user)
            else:
                self.logger.error(f"Failed to accept user {user}.")
            time.sleep(self.RATE_LIMIT)

        if users_accepted != users_to_accept:
            self.logger.warning(f"Some users could not be accepted: {list(users_to_accept - users_accepted)}")

        # 4. Finalize all users in this list (RATE LIMITED)
        for user in users_accepted:
            if self.finalize_user(user):
                users_finalized.add(user)
            else:
                self.logger.error(f"Failed to finalize user {user}.")
            time.sleep(self.RATE_LIMIT)

        if users_finalized != users_accepted:
            self.logger.warning(f"Some users could not be finalized: {list(users_accepted - users_finalized)}")

        # 5. Send email with the user ids and template (POST /mail/send/)
        if self.send_email(list(users_finalized), "acceptance"):
            self.logger.info(f"Successfully sent mail to users: {list(users_finalized)}.")
        else:
            self.logger.error(f"Failed to send mail to users: {list(users_finalized)}.")

    def run_in_loop(self) -> None:
        while True:
            self.logger.info("=====Starting accept cycle=====")
            try:
                bot.run_accept_cycle()
            except Exception as e:
                self.logger.error(f"Error in run_accept_cycle: {e}")
            self.logger.info("=====Finished accept cycle=====")

            self.logger.info(f"Sleeping for {SLEEP_INTERVAL} seconds.")
            time.sleep(SLEEP_INTERVAL)

if __name__ == "__main__":
    admin_jwt = os.environ["HackIllinois_Admin_JWT"]
    bot = AutoAccept(admin_jwt)
    bot.run_in_loop()
