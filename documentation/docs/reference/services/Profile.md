Profile
============

!!! warning
    The `id` in the profile service refers to a separate, randomly-generated, profile-only id. 
    **This is different from the user `id` used in other services.**
    When a profile is created, a mapping from the user `id` to the profile `id` is stored in the database.
    
    We will distinguish user id and profile id by using `github123456` and `profileid123456` for each, respectively, in the examples below.

GET /profile/
-------------------------

Returns the profile stored for the current user.

Valid values for `teamStatus` are `LOOKING_FOR_MEMBERS`, `LOOKING_FOR_TEAM`, and `NOT_LOOKING`.

Request requires no body.

```json title="Example response"
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234"
}

```

GET /profile/ID/
-------------------------

Returns the profile stored for user that has the `id` `ID`.

Request requires no body.

```json title="Example response"
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234"
}
```

GET /profile/list/?teamStatus=value&interests=value,value,value&limit=value
-------------------------

!!! danger
    Internal use only!
    If you are looking to search for profiles as an attendee, applicant, or mentor, please use
    [GET /profile/search/](#get-profilesearchteamstatusvalueinterestsvaluevaluevaluelimitvalue).

Returns a list of profiles matching the filter conditions.

`teamStatus` is a string matching the user's team status.

`interests` is a comma-separated string representing the user's interests.

- i.e if the user's interests are ["C++", "Machine Learning"], you can filter on this by sending `interests="C++,Machine Learning"`

If a `limit` parameter is provided, it will return the first matching `limit` profiles. Otherwise, it will return all of the matched profiles.

If no parameters are provided, it returns all profiles that are in the database.

Request requires no body.

```json title="Example response"
{
    "profiles": [
        {
            "id": "profileid123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "timezone": "Americas UTC+8",
            "avatarUrl": "https://github.com/.../profile.jpg",
            "discord": "patrick#1234"
        },
        {
            "id": "profileid123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "timezone": "Americas UTC+8",
            "avatarUrl": "https://github.com/.../profile.jpg",
            "discord": "patrick#1234"
        },
    ]
}
```

POST /profile/
-------------------

Creates a profile for the currently authenticated user (determined by the JWT in the `Authorization` header).

```json title="Example request"
{
    "firstName": "John",
    "lastName": "Doe",
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234"
}
```

```json title="Example response"
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234"
}
```

PUT /profile/
------------------

Updates the profile for the currently authenticated user (determined by the JWT in the `Authorization` header).

!!! warning
    You can not edit the `points` field through this (for security reasons)

```json title="Example request"
{
    "firstName": "John",
    "lastName": "Doe",
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234"
}
```

```json title="Example response"
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234"
}
```


DELETE /profile/
------------------

!!! danger
    Temporarily disabled since Apr 2nd, 2021.

Deletes the profile for the currently authenticated user (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234"
}
```

GET /profile/leaderboard/?limit=
-------------------------

!!! note
    This is a public endpoint

Returns a list of profiles sorted by points descending. If a `limit` parameter is provided, it will return the first `limit` profiles. Otherwise, it will return all of the profiles.

Request requires no body.

```json title="Example response"
{
    "profiles": [
        {
            "id": "profileid123456",
            "points": 2021,
            "discord": "patrick#1234"
        },
        {
            "id": "profileid123456",
            "points": 2021,
            "discord": "patrick#1234"
        },
    ]
}
```

WS /profile/live/leaderboard/?limit=
-------------------------

!!! note
    This is a public endpoint

This endpoint expects a websocket connection.

Returns a list of profiles sorted by points descending. If a `limit` parameter is provided, it will
return the first `limit` profiles. Otherwise, it will return all of the profiles.

Once, the connection has been established, `limit` can be easily changed by simply sending the key
in a JSON.

The API will send back a list of profiles whenever a user gets awarded points (ratelimited to 1
response per second) or if 5 minutes elapse with no activity (i.e. no point values change). The
latter is more or less a failsafe in order to ensure the leaderboard is the most up-to-date.

```json title="Exmaple WS message request"
{
    "limit": 20
}
```

```json title="Example WS response"
{
    "profiles": [
        {
            "id": "profileid123456",
            "points": 2021,
            "discord": "patrick#1234"
        },
        {
            "id": "profileid123456",
            "points": 2021,
            "discord": "patrick#1234"
        },
    ]
}
```

GET /profile/search/?teamStatus=value&interests=value,value,value&limit=value
-------------------------

Returns a list of profiles matching the filter conditions. 

`teamStatus` is a string matching the user's team status. Valid values for `teamStatus` are `LOOKING_FOR_MEMBERS`, `LOOKING_FOR_TEAM`, and `NOT_LOOKING`.

`interests` is a comma-separated string representing the user's interests.

- i.e if the user's interests are ["C++", "Machine Learning"], you can filter on this by sending `interests="C++,Machine Learning"`

If a `limit` parameter is provided, it will return the first matching `limit` profiles. Otherwise, it will return all of the matched profiles.

!!! warning
    Users with `teamStatus` `NOT_LOOKING` will be shown unless you filter against it.
    **Make sure you consider this** in respect to your use case and act accordingly!

Request requires no body.

```json title="Example response"
{
    "profiles": [
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

POST /profile/event/checkin/
----------------------------

!!! danger
    Internal use only!

Validates the status of an event that the user is trying to check into.
This is an internal endpoint hit during the checkin process (when the user posts a code to the event service).
The response is a status string, and throws an error (except the case when the user is already checked in).
In the case that the user has already been checked, status is set to "Event already redeemed" and a 200 status code is still used.

!!! note
    Here, the "id" actually refers to the user id, not the profile id (hence `github123456` instead of `profileid123456`)

```json title="Example request"
{
    "id": "github123456",
    "eventID": "52fdfc072182654f163f5f0f9a621d72"
}

```

```json title="Example response"
{
    "status": "Success"
}
```

POST /profile/points/award/
----------------------------

!!! danger
    Internal use only!

Takes a struct with a profile and a certain number of points to increment their score by, and returns this profile upon completion.

Note: here, the "id" actually refers to the user id, not the profile id (hence `github123456` instead of `profileid123456`)

```json title="Example request"
{
    "id": "github123456",
    "points": 10
}

```

```json title="Example response"
{
    "id": "profileid123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234"
    "points": 10
}
```

GET /profile/favorite/
-------------------------

Returns a list of profiles that the current user has favorited.

Request requires no body.

```json title="Example response"
{
    "id": "testid", 
    "profiles": [
        "testid3",
    ]
}
```

POST /profile/favorite/
-------------------------

Adds the specified profile to the current user's favorite list, and returns the updated list of favorite profiles.

```json title="Example request"
{
    "id": "testid2"
}
```

```json title="Example response"
{
    "id": "testid",
    "profiles": [
        "testid3",
        "testid2"
    ]
}
```

DELETE /profile/favorite/
-------------------------

Removes the specified profile from the current user's favorite list, and returns the updated list of favorite profiles.

```json title="Example request"
{
    "id": "testid3"
}
```

```json title="Example response"
{
    "id": "testid",
    "profiles": [
        "testid2"
    ]
}
```

GET /profile/tier/threshold/
-------------------------
Returns the profile tier name to minimum point threshold mapping.

Request requires no body.

```json title="Example response"
[
  {
    "name": "cookie",
    "threshold": 0
  },
  {
    "name": "cupcake",
    "threshold": 50
  }
]
```
