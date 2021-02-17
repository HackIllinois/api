Profile
============


GET /profile/
-------------------------

Returns the profile stored for the current user.

Response format:
```
{
    "id": "github0000001"
    "name": "John Smith",
    "email": "john@gmail.com",
    "github": "JSmith",
    "linkedin": "john-smith",
    "interests": [
        "deep learning",
        "python"
    ]
}
```

GET /profile/{id}
-------------------------

Returns the profile stored for user that has the ID ``{id}``.

Response format:
```
{
    "id": "github0000001"
    "name": "John Smith",
    "email": "john@gmail.com",
    "github": "JSmith",
    "linkedin": "john-smith",
    "interests": [
        "deep learning",
        "python"
    ]
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
            "id": "github0000001"
            "name": "John Smith",
            "email": "john@gmail.com",
            "github": "JSmith",
            "linkedin": "john-smith",
            "interests": [
                "deep learning",
                "ice skating",
                "python"
            ]
        },
        {
            "id": "github0000002"
            "name": "Smith John",
            "email": "smith@gmail.com",
            "github": "SJohn",
            "linkedin": "smith-john",
            "interests": [
                "classical music",
                "deep learning",
                "python"
            ]
        }
```


POST /profile/
-------------------

Creates a profile for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
    "name": "John Smith",
    "email": "john@gmail.com",
    "github": "JSmith",
    "linkedin": "john-smith",
    "interests": [
        "python",
        "deepLearning"
    ]
}
```

Response format:
```
{
    "id": "github0000001"
    "name": "John Smith",
    "email": "john@gmail.com",
    "github": "JSmith",
    "linkedin": "john-smith",
    "interests": [
        "deep learning",
        "python"
    ]
}
```

PUT /profile/
------------------

Updates the profile for the user with the `id` in the JWT token provided in the Authorization header.
This returns the updated profile information.

Request format:
```
{
    "name": "John Smith",
    "email": "john@gmail.com",
    "github": "JSmith",
    "linkedin": "john-smith",
    "interests": [
        "deep learning",
        "python"
    ]
}
```

Response format:
```
{
    "id": "github0000001"
    "name": "John Smith",
    "email": "john@gmail.com",
    "github": "JSmith",
    "linkedin": "john-smith",
    "interests": [
        "deep learning",
        "python"
    ]
}
```


DELETE /profile/
------------------

Deletes the profile for the user with the `id` in the JWT token provided in the Authorization header.
This returns the deleted profile information.

Response format:
```
{
    "id": "github0000001"
    "name": "John Smith",
    "email": "john@gmail.com",
    "github": "JSmith",
    "linkedin": "john-smith",
    "interests": [
        "deep learning",
        "python"
    ]
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
            "id": "github0000001"
            "name": "John Smith",
            "email": "john@gmail.com",
            "github": "JSmith",
            "linkedin": "john-smith",
            "interests": [
                "deep learning",
                "ice skating",
                "python"
            ],
            "points": 1000,
        },
        {
            "id": "github0000002"
            "name": "Smith John",
            "email": "smith@gmail.com",
            "github": "SJohn",
            "linkedin": "smith-john",
            "interests": [
                "classical music",
                "deep learning",
                "python"
            ],
            "points": 500,
        }
```

GET /profile/search/?teamStatus=value&interests=value,value,value&limit=value
-------------------------

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
            "id": "github0000001"
            "name": "John Smith",
            "email": "john@gmail.com",
            "github": "JSmith",
            "linkedin": "john-smith",
            "interests": [
                "deep learning",
                "ice skating",
                "python"
            ],
            "points": 1000,
        },
        {
            "id": "github0000002"
            "name": "Smith John",
            "email": "smith@gmail.com",
            "github": "SJohn",
            "linkedin": "smith-john",
            "interests": [
                "classical music",
                "deep learning",
                "python"
            ],
            "points": 500,
        }
```
