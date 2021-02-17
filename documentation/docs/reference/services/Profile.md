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