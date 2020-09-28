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
        "python",
        "deepLearning"
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
        "python",
        "deepLearning"
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
        "python",
        "deepLearning"
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
        "python",
        "deepLearning"
    ]
}
```