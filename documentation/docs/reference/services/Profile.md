Profile
============

GET /profile/
-------------------------

Returns the profile stored for the current user.

***The `id` in the profile service refers to a separate, randomly-generated, profile-only id. This is different from the (user) `id` used in other services. When a profile is created, a mapping from the user `id` to the profile `id` is stored in the database.***

We will distinguish user id and profile id by using `github123456` and `profileid123456` for each, respectively, in the examples below.

Response format:
```
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "isVirtual": false
}

```

GET /profile/{id}/
-------------------------

Returns the profile stored for user that has the profile ID ``{id}``.

Response format:
```
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "isVirtual": false
}
```

GET /profile/user/{id}/
-------------------------

**Internal use only.**

Returns the profile stored for user that has the user ID ``{id}``.

Response format:
```
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "isVirtual": false
}
```

GET /profile/list/?isVirtual=value&interests=value,value,value&limit=value
-------------------------

**Internal use only.**

Returns a list of profiles matching the filter conditions.

teamStatus is a string matching the user's team status.

interests is a comma-separated string representing the user's interests.

If a ``limit`` parameter is provided, it will return the first matching ``limit`` profiles. Otherwise, it will return all of the matched profiles.

If no parameters are provided, it returns all profiles that are in the database.

Response format:
```
{
    profiles: [
        {
            "id": "profileid123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "isVirtual": false
        },
        {
            "id": "profileid123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "isVirtual": false
        },
    ]
}
```

POST /profile/
-------------------

Creates a profile for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
    "firstName": "John",
    "lastName": "Doe",
    "isVirtual": false
}
```

Response format:
```
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 0,
    "isVirtual": false
}
```

PUT /profile/
------------------

Updates the profile for the user with the `id` in the JWT token provided in the Authorization header.
This returns the updated profile information.
Note you can not edit the ``points`` field through this.

Request format:
```
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "isVirtual": true
}
```

Response format:
```
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 0,
    "isVirtual": true
}
```


DELETE /profile/
------------------

**Temporarily disabled**

Deletes the profile for the user with the `id` in the JWT token provided in the Authorization header.
This returns the deleted profile information.

Response format:
```
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 0,
    "isVirtual": false
}
```

GET /profile/leaderboard/?limit=
-------------------------

**Public endpoint.**

Returns a list of profiles sorted by points descending. If a ``limit`` parameter is provided, it will return the first ``limit`` profiles. Otherwise, it will return all of the profiles.

Response format:
```
{
    profiles: [
        {
            "id": "profileid123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
        },
        {
            "id": "profileid123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
        },
    ]
}
```

GET /profile/search/?isVirtual=value&interests=value,value,value&limit=value
-------------------------

Returns a list of profiles matching the filter conditions.

If a ``limit`` parameter is provided, it will return the first matching ``limit`` profiles. Otherwise, it will return all of the matched profiles.

Response format:
```
{
    profiles: [
        {
            "id": "profileid123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "isVirtual": false
        },
        {
            "id": "profileid123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "isVirtual": false
        },
    ]
}
```

POST /profile/event/checkin/
----------------------------

**Internal Use Only**
Validates the status of an event that the user is trying to check into.
This is an internal endpoint hit during the checkin process (when the user posts a code to the event service).
The response is a status string, and throws an error (except the case when the user is already checked in).
In the case that the user has already been checked, status is set to "Event already redeemed" and a 200 status code is still used.

Note: here, the "id" actually refers to the user id, not the profile id (hence `github123456` instead of `profileid123456`)

Request format:
```
{
    "id": "github123456",
    "eventID": "52fdfc072182654f163f5f0f9a621d72"
}

```

Response format:
```
{
    "status": "Success"
}
```

POST /profile/points/award/
----------------------------

**Internal Use Only**
Takes a struct with a profile and a certain number of points to increment their score by, and returns this profile upon completion.

Note: here, the "id" actually refers to the user id, not the profile id (hence `github123456` instead of `profileid123456`)

Request format:
```
{
    "id": "github123456",
    "points": 10
}

```

Response format:
```
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 10,
    "isVirtual": false
}
```

GET /profile/favorite/
-------------------------

Returns a list of profiles that the current user has favorited.

Response format:
```
{
    id: "testid", 
    profiles: [
        "testid3",
    ]
}
```

POST /profile/favorite/
-------------------------
Adds the specified profile to the current user's favorite list, and returns the updated list of favorite profiles.

Request format:
```
{
    id: "testid2"
}
```

Response format:
```
{
    id: "testid"
    profiles: [
        "testid3",
        "testid2",
    ]
}
```

DELETE /profile/favorite/
-------------------------
Removes the specified profile from the current user's favorite list, and returns the updated list of favorite profiles.

Request format:
```
{
    id: "testid3"
}
```

Response format:
```
{
    id: "testid"
    profiles: [
        "testid2",
    ]
}
```
