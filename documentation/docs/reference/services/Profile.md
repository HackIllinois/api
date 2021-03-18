Profile
============


GET /profile/
-------------------------

Returns the profile stored for the current user.

Response format:
```
{
    "id": "github123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234",
    "teamStatus": "looking",
    "description": "Lorem Ipsum…",
    "interests": ["C++", "Machine Learning"]
}

```

GET /profile/{id}
-------------------------

Returns the profile stored for user that has the ID ``{id}``.

Response format:
```
{
    "id": "github123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234",
    "teamStatus": "looking",
    "description": "Lorem Ipsum…",
    "interests": ["C++", "Machine Learning"]
}
```

GET /profile/list/
-------------------------

Returns all profiles that are in the database.

Response format:
```
{
    profiles: [
        {
            "id": "github123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "timezone": "Americas UTC+8",
            "avatarUrl": "https://github.com/.../profile.jpg",
            "discord": "patrick#1234",
            "teamStatus": "looking",
            "description": "Lorem Ipsum…",
            "interests": ["C++", "Machine Learning"]
        },
        {
            "id": "github123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "timezone": "Americas UTC+8",
            "avatarUrl": "https://github.com/.../profile.jpg",
            "discord": "patrick#1234",
            "teamStatus": "looking",
            "description": "Lorem Ipsum…",
            "interests": ["C++", "Machine Learning"]
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
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234",
    "teamStatus": "looking",
    "description": "Lorem Ipsum…",
    "interests": ["C++", "Machine Learning"]
}
```

Response format:
```
{
    "id": "github123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234",
    "teamStatus": "looking",
    "description": "Lorem Ipsum…",
    "interests": ["C++", "Machine Learning"]
}
```

PUT /profile/
------------------

Updates the profile for the user with the `id` in the JWT token provided in the Authorization header.
This returns the updated profile information.

Request format:
```
{
    "firstName": "John",
    "lastName": "Doe",
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234",
    "teamStatus": "looking",
    "description": "Lorem Ipsum…",
    "interests": ["C++", "Machine Learning"]
}
```

Response format:
```
{
    "id": "github123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234",
    "teamStatus": "looking",
    "description": "Lorem Ipsum…",
    "interests": ["C++", "Machine Learning"]
}
```


DELETE /profile/
------------------

Deletes the profile for the user with the `id` in the JWT token provided in the Authorization header.
This returns the deleted profile information.

Response format:
```
{
    "id": "github123456",
    "firstName": "John",
    "lastName": "Doe",
    "points": 2021,
    "timezone": "Americas UTC+8",
    "avatarUrl": "https://github.com/.../profile.jpg",
    "discord": "patrick#1234",
    "teamStatus": "looking",
    "description": "Lorem Ipsum…",
    "interests": ["C++", "Machine Learning"]
}
```

GET /profile/leaderboard/?limit=
-------------------------

Returns a list of profiles sorted by points descending. If a ``limit`` parameter is provided, it will return the first ``limit`` profiles. Otherwise, it will return all of the profiles.

Response format:
```
{
    profiles: [
        {
            "id": "github123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
        },
        {
            "id": "github123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
        },
    ]
}
```

GET /profile/search/?teamStatus=value&interests=value,value,value&limit=value
-------------------------

Returns a list of profiles matching the filter conditions. 

teamStatus is a string matching the user's team status.

interests is a comma-separated string representing the user's interests.

- i.e if the user's interests are ["C++", "Machine Learning"], you can filter on this by sending ``interests="C++,Machine Learning"``

If a ``limit`` parameter is provided, it will return the first matching ``limit`` profiles. Otherwise, it will return all of the matched profiles.

Any users with the TeamStatus "Not Looking" will be removed.

Response format:
```
{
    profiles: [
        {
            "id": "github123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "timezone": "Americas UTC+8",
            "avatarUrl": "https://github.com/.../profile.jpg",
            "discord": "patrick#1234",
            "teamStatus": "looking",
            "description": "Lorem Ipsum…",
            "interests": ["C++", "Machine Learning"]
        },
        {
            "id": "github123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "timezone": "Americas UTC+8",
            "avatarUrl": "https://github.com/.../profile.jpg",
            "discord": "patrick#1234",
            "teamStatus": "looking",
            "description": "Lorem Ipsum…",
            "interests": ["C++", "Machine Learning"]
        },
    ]
}
```

GET /profile/filtered/?teamStatus=value&interests=value,value,value&limit=value
-------------------------

**Internal use only.**

Returns a list of profiles matching the filter conditions. 

teamStatus is a string matching the user's team status.

interests is a comma-separated string representing the user's interests.

- i.e if the user's interests are ["C++", "Machine Learning"], you can filter on this by sending ``interests="C++,Machine Learning"``

If a ``limit`` parameter is provided, it will return the first matching ``limit`` profiles. Otherwise, it will return all of the matched profiles.

Response format:
```
{
    profiles: [
        {
            "id": "github123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "timezone": "Americas UTC+8",
            "avatarUrl": "https://github.com/.../profile.jpg",
            "discord": "patrick#1234",
            "teamStatus": "looking",
            "description": "Lorem Ipsum…",
            "interests": ["C++", "Machine Learning"]
        },
        {
            "id": "github123456",
            "firstName": "John",
            "lastName": "Doe",
            "points": 2021,
            "timezone": "Americas UTC+8",
            "avatarUrl": "https://github.com/.../profile.jpg",
            "discord": "patrick#1234",
            "teamStatus": "Not Looking",
            "description": "Lorem Ipsum…",
            "interests": ["C++", "Machine Learning"]
        },
    ]
}
```

