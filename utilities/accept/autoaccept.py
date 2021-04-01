import requests
import time
import logging
import os

class AutoAccept():
    def __init__(self, admin_jwt: str):
        self.headers = {
            "Content-Type": "application/json",
            "Authorization": admin_jwt,
        }
        self.logger = logging.Logger("AutoAccept")
        self.logger.addHandler(logging.FileHandler("./AutoAccept.log"))

        self.RATE_LIMIT = 0.10  # time to wait between rate-limit-sensitive API requests, in seconds
        self.base_url = "https://api.hackillinois.org"

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
            Accepts a user to the event with ``id``. Returns a bool representing if the accept was successful. 
        '''

        data = {
            "id": id,
            "status": "ACCEPTED",
            "wave": 1
        }
        r = requests.post(f"{self.base_url}/decision/", headers=self.headers, data=data)
        return r.ok

    def finalize_user(self, id: str) -> bool:
        '''
            Finalizes a user's decision. Returns a bool representing if the finalization was successful.
        '''

        data = {
            "id": id,
            "finalized": True
        }

        r = requests.post(f"{self.base_url}/decision/finalize/", headers=self.headers, data=data)
        return r.ok

    def send_email(self, ids: list, template_id: str) -> bool:
        data = {
            "ids": ids,
            "template": template_id
        }

        r = requests.post(f"{self.base_url}/mail/send/", headers=self.headers, data=data)
        return r.ok

    def run_accept_cycle(self):
        # 1. Get a list of all registrations (entire pool of Users that have registered)
        registered_users = self.get_registered_users()
        
        # 2. Subtract the list of users that have already been accepted (accepted_wave_1 Mail List)
        accepted_users = self.get_accepted_users()

        users_to_accept = registered_users - accepted_users
        users_to_accept = users_to_accept[0]

        # 3. Accept all users in this list (RATE LIMITED)
        for user in users_to_accept:
            if not self.accept_user(user):
                self.logger.error(f"Failed to accept user {user}.")
            time.sleep(self.RATE_LIMIT)

        # 4. Finalize all users in this list (RATE LIMITED)
        for user in users_to_accept:
            if not self.finalize_user(user):
                self.logger.error(f"Failed to finalize user {user}.")
            time.sleep(self.RATE_LIMIT)

        # 5. Send email with the user ids and template (POST /mail/send/)
        self.send_email(list(users_to_accept), "TODO_TEMPLATE_ID")

if __name__ == "__main__":
    admin_jwt = os.environ["HackIllinois_Admin_JWT"]
    bot = AutoAccept(admin_jwt)
    bot.run_accept_cycle()
