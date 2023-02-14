import os
import time
import requests
import sys

ADMIN_JWT_ENV_NAME = "HackIllinois_Admin_JWT"
BASE_URL_ENV_NAME = "HackIllinois_Base_Url"

RATE_LIMIT = 0.10  # Time to wait between each individual request, DO NOT SET TO 0


def diet_exists(diet):
    if diet:
        if type(diet) is str:
            if diet.lower() == "none":
                return True
        elif type(diet) is list:
            for item in diet:
                if item.lower() != "none":
                    return True

    return False


class FoodWave:
    def __init__(self, waves: int, admin_jwt: str, base_url: str):
        self.headers = {
            "Content-Type": "application/json",
            "Authorization": admin_jwt,
        }

        self.base_url = base_url
        self.waves = waves

    def __get_all_users(self):
        # Filter returns all users that match its criteria.
        # Kind of weird we don't have a defined way to get all users other than this.
        # Probably should be implemented, but this works™️
        response = requests.get(f"{self.base_url}/user/filter/", headers=self.headers)
        response.raise_for_status()

        # AOC prepared me for this.
        # This takes the response json, takes the users, and maps them to their id.
        # Then, it stores the ids in a set.
        return set(map(lambda x: x["id"], response.json()["users"]))

    def __get_rsvp_users(self):
        users = list(self.__get_all_users())

        rsvp_users = dict()
        last_notice = None

        for i in range(len(users)):
            id = users[i]
            response = requests.get(f"{base_url}/rsvp/{id}/", headers=self.headers)

            if response.status_code == 200:
                json = response.json()
                if json["isAttending"] == True:
                    rsvp_users[id] = json

            time.sleep(RATE_LIMIT)

            notice_percent = i / len(users)

            if not last_notice or notice_percent >= last_notice + 0.1:
                last_notice = notice_percent
                print(f"Fetching RSVP data, {notice_percent:.0%}")

        return rsvp_users

    def __get_rsvp_users_to_has_dietary(self):
        rsvp_users = self.__get_rsvp_users()
        rsvp_users_to_has_dietary = dict()

        for [id, user_info] in rsvp_users.items():
            rsvp_users_to_has_dietary[id] = diet_exists(user_info["diet"])

        return rsvp_users_to_has_dietary

    def __get_rsvp_users_sorted_by_priority(self):
        rsvp_users_to_has_dietary = self.__get_rsvp_users_to_has_dietary()

        sorted_users = list(
            sorted(rsvp_users_to_has_dietary.items(), key=lambda x: x[1], reverse=True)
        )

        return list(map(lambda x: x[0], sorted_users))

    def assign_food_waves(self):
        rsvp_users_sorted_by_priority = self.__get_rsvp_users_sorted_by_priority()

        last_notice = None
        assigned = 0

        for i in range(len(rsvp_users_sorted_by_priority)):
            id = rsvp_users_sorted_by_priority[i]
            wave = int((i / len(rsvp_users_sorted_by_priority)) * waves) + 1

            self.headers["HackIllinois-Impersonation"] = id

            response = requests.get(
                f"{base_url}/profile/",
                headers=self.headers,
            )

            if response.status_code == 200:
                profile = response.json()

                profile["foodWave"] = wave

                response = requests.put(
                    f"{base_url}/profile/",
                    json=profile,
                    headers=self.headers,
                )

                if response.status_code == 200:
                    assigned += 1

            del self.headers["HackIllinois-Impersonation"]

            time.sleep(RATE_LIMIT)

            notice_percent = i / len(rsvp_users_sorted_by_priority)
            if not last_notice or notice_percent >= last_notice + 0.1:
                last_notice = notice_percent
                print(f"Assigning food waves, {notice_percent:.0%}")

        print(
            f"Assigned food waves to {assigned}/{len(rsvp_users_sorted_by_priority)} users"
        )


if __name__ == "__main__":
    admin_jwt = os.environ.get(ADMIN_JWT_ENV_NAME)

    if len(sys.argv) != 2:
        raise Exception("Proper usage: python foodwave.py [WAVE]")

    waves = int(sys.argv[1])
    print(f"Splitting users across {waves} waves")

    if not admin_jwt:
        raise Exception(f"Please set the `{ADMIN_JWT_ENV_NAME}` environment variable")

    print(f"Loaded admin JWT from `{ADMIN_JWT_ENV_NAME}` environment variable")

    base_url = os.environ.get(BASE_URL_ENV_NAME) or "https://api.hackillinois.org"

    if "http://" not in base_url and "https://" not in base_url:
        raise Exception(
            f"The base url ({base_url}) must include a protocol (http://, https://). Set the `{BASE_URL_ENV_NAME}` to change this."
        )

    print(
        f"Using base url `{base_url}`. You can set the `{BASE_URL_ENV_NAME}` to change this."
    )

    client = FoodWave(waves, admin_jwt, base_url)
    client.assign_food_waves()
