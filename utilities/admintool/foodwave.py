import os
import time
from typing import Any
import requests
import sys

ADMIN_JWT_ENV_NAME = "HackIllinois_Admin_JWT"
BASE_URL_ENV_NAME = "HackIllinois_Base_Url"

RATE_LIMIT = 0.10  # Time to wait between each individual request, DO NOT SET TO 0


# Returns true if the diet passed represents dietary restrictions
# This is a thing because "None" was an option...
def diet_exists(diet) -> bool:
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

    # Gets all the users, returned as a set of ids
    def __get_all_users(self) -> "set[str]":
        # Filter returns all users that match its criteria.
        # Kind of weird we don't have a defined way to get all users other than this.
        # Probably should be implemented, but this works™️
        response = requests.get(f"{self.base_url}/user/filter/", headers=self.headers)
        response.raise_for_status()

        # AOC prepared me for this.
        # This takes the response json, takes the users, and maps them to their id.
        # Then, it stores the ids in a set.
        return set(map(lambda x: x["id"], response.json()["users"]))

    # Gets all the users, that are RSVP'd, returned as a dict of id to RSVP data
    def __get_rsvp_users(self) -> "dict[str, Any]":
        # First, get all users
        users = list(self.__get_all_users())

        rsvp_users = dict()
        last_notice = None

        # Loop through all the users, requesting their RSVP data
        for i in range(len(users)):
            id = users[i]
            response = requests.get(f"{base_url}/rsvp/{id}/", headers=self.headers)

            # If they're attending, add them to the list of results
            if response.status_code == 200:
                json = response.json()
                if json["isAttending"] == True:
                    rsvp_users[id] = json

            # Rate limit
            time.sleep(RATE_LIMIT)

            # Notice output
            notice_percent = i / len(users)

            if not last_notice or notice_percent >= last_notice + 0.1:
                last_notice = notice_percent
                print(f"Fetching RSVP data, {notice_percent:.0%}")

        return rsvp_users

    # Gets all the users that are RSVP'd, and returns a dict mapping each id to wether or not they have dietary restrictions
    def __get_rsvp_users_to_has_dietary(self) -> "dict[str, bool]":
        rsvp_users = self.__get_rsvp_users()
        rsvp_users_to_has_dietary = dict()

        for [id, user_info] in rsvp_users.items():
            rsvp_users_to_has_dietary[id] = diet_exists(user_info["diet"])

        return rsvp_users_to_has_dietary

    # Gets all the users that are RSVP'd, and sorts them by priority. Higher priority = further up, dietary restrictions = priority
    def __get_rsvp_users_sorted_by_priority(self):
        rsvp_users_to_has_dietary = self.__get_rsvp_users_to_has_dietary()

        sorted_users = list(
            sorted(rsvp_users_to_has_dietary.items(), key=lambda x: x[1], reverse=True)
        )

        return list(map(lambda x: x[0], sorted_users))

    # Gets all the users that are RSVP'd sorted by priority, and assigns food waves accordingly.
    # Highest priority gets first wave, lowest priority gets last wave
    def assign_food_waves(self):
        rsvp_users_sorted_by_priority = self.__get_rsvp_users_sorted_by_priority()

        last_notice = None
        assigned = 0

        # Go through each user, assigning a food wave
        for i in range(len(rsvp_users_sorted_by_priority)):
            id = rsvp_users_sorted_by_priority[i]

            # Earliest should get lowest wave, so we do fancy maths to assign waves
            wave = int((i / len(rsvp_users_sorted_by_priority)) * waves) + 1

            self.headers["HackIllinois-Impersonation"] = id

            # We need to first GET their current profile to see if it exists
            response = requests.get(
                f"{base_url}/profile/",
                headers=self.headers,
            )

            # Then, if it does, we need to PUT the same data but with a modified food wave
            if response.status_code == 200:
                profile = response.json()

                profile["foodWave"] = wave

                response = requests.put(
                    f"{base_url}/profile/",
                    json=profile,
                    headers=self.headers,
                )

                # If that all works, we can increment our success counter
                if response.status_code == 200:
                    assigned += 1

            del self.headers["HackIllinois-Impersonation"]

            # Rate limit
            time.sleep(RATE_LIMIT)

            # Notice output
            notice_percent = i / len(rsvp_users_sorted_by_priority)
            if not last_notice or notice_percent >= last_notice + 0.1:
                last_notice = notice_percent
                print(f"Assigning food waves, {notice_percent:.0%}")

        print(
            f"Assigned food waves to {assigned}/{len(rsvp_users_sorted_by_priority)} users"
        )


# Only runs when called directly, otherwise we export the class
if __name__ == "__main__":
    # Input validation
    if len(sys.argv) != 2:
        raise Exception("Proper usage: python foodwave.py [WAVE]")

    waves = int(sys.argv[1])
    print(f"Splitting users across {waves} waves")

    admin_jwt = os.environ.get(ADMIN_JWT_ENV_NAME)

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

    # Run
    client = FoodWave(waves, admin_jwt, base_url)
    client.assign_food_waves()
