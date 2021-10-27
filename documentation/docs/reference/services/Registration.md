Registration
============

*Note:* The exact fields in the registration requests and responses will change based on the registration definitions provided in the API configuration file.

GET /registration/
-------------------------

Returns all the registrations stored for the current user. If registrations are not found for either Attendee or Mentor,
that field is set to null.

Response format:
```
{
    "attendee": {
        "createdAt": 1635290385,
        "degreePursued": "",
        "email": "john@gmail.com",
        "firstName": "John",
        "gender": "MALE",
        "github": "localadmin",
        "graduationYear": 2019,
        "hasInternship": false,
        "id": "localadmin",
        "lastName": "Smith",
        "location": "Champaign",
        "major": "Computer Science",
        "programmingAbility": 4,
        "programmingYears": 0,
        "race": null,
        "resumeFilename": "",
        "school": "University of Illinois at Urbana-Champaign",
        "timezone": "Central",
        "updatedAt": 1635290385
    },
    "mentor": {
        "createdAt": 1635290586,
        "email": "john@gmail.com",
        "firstName": "John",
        "github": "localadmin",
        "id": "localadmin",
        "lastName": "Smith",
        "linkedin": "john-smith",
        "shirtSize": "M",
        "updatedAt": 1635290586
    }
}
```

GET /registration/USERID/
-------------------------

Returns all registrations stored for the user with the `id` `USERID`.
If registrations are not found for either Attendee or Mentor, that field is set to null.

Response format:
```
{
    "attendee": {
        "createdAt": 1635290385,
        "degreePursued": "",
        "email": "john@gmail.com",
        "firstName": "John",
        "gender": "MALE",
        "github": "localadmin",
        "graduationYear": 2019,
        "hasInternship": false,
        "id": "localadmin",
        "lastName": "Smith",
        "location": "Champaign",
        "major": "Computer Science",
        "programmingAbility": 4,
        "programmingYears": 0,
        "race": null,
        "resumeFilename": "",
        "school": "University of Illinois at Urbana-Champaign",
        "timezone": "Central",
        "updatedAt": 1635290385
    },
    "mentor": {
        "createdAt": 1635290586,
        "email": "john@gmail.com",
        "firstName": "John",
        "github": "localadmin",
        "id": "localadmin",
        "lastName": "Smith",
        "linkedin": "john-smith",
        "shirtSize": "M",
        "updatedAt": 1635290586
    }
}
```

GET /registration/attendee/USERID/
-------------------------

Returns the user registration stored for the user with the `id` `USERID`.

Response format:
```
{
    "createdAt": 1635290385,
    "degreePursued": "",
    "email": "john@gmail.com",
    "firstName": "John",
    "gender": "MALE",
    "github": "localadmin",
    "graduationYear": 2019,
    "hasInternship": false,
    "id": "localadmin",
    "lastName": "Smith",
    "location": "Champaign",
    "major": "Computer Science",
    "programmingAbility": 4,
    "programmingYears": 0,
    "race": null,
    "resumeFilename": "",
    "school": "University of Illinois at Urbana-Champaign",
    "timezone": "Central",
    "updatedAt": 1635290385
}
```

GET /registration/attendee/
------------------

Returns the user registration stored for the user with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
    "createdAt": 1635290385,
    "degreePursued": "",
    "email": "john@gmail.com",
    "firstName": "John",
    "gender": "MALE",
    "github": "localadmin",
    "graduationYear": 2019,
    "hasInternship": false,
    "id": "localadmin",
    "lastName": "Smith",
    "location": "Champaign",
    "major": "Computer Science",
    "programmingAbility": 4,
    "programmingYears": 0,
    "race": null,
    "resumeFilename": "",
    "school": "University of Illinois at Urbana-Champaign",
    "timezone": "Central",
    "updatedAt": 1635290385
}
```

POST /registration/attendee/
-------------------

Creates a registration for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
    "firstName": "John",
    "lastName": "Smith",
    "email": "john@gmail.com",
    "shirtSize": "M",
    "diet": "NONE",
    "age": 19,
    "graduationYear": 2019,
    "transportation": "NONE",
    "school": "University of Illinois at Urbana-Champaign",
    "major": "Computer Science",
    "gender": "MALE",
    "professionalInterest": "INTERNSHIP",
    "github": "JSmith",
    "linkedin": "john-smith",
    "location": "Champaign",
    "interests": "Software",
    "isNovice": false,
    "isPrivate": false,
    "phoneNumber": "555-555-5555",
    "programmingAbility": 4,
    "timezone": "Central",
    "longforms": [
        {
            "response": "This is a longform."
        }
    ],
    "extraInfos": [
        {
            "response": "This is an extra info."
        }
    ],
    "osContributors": [
        {
            "name": "Tom",
            "contactInfo": "tom@gmail.com"
        }
    ],
    "collaborators": [
        {
            "github": "collabgithub"
        }
    ]
}
```

Response format:
```
{
    "createdAt": 1635291098,
    "degreePursued": "",
    "email": "john@gmail.com",
    "firstName": "John",
    "gender": "MALE",
    "github": "localadmin",
    "graduationYear": 2019,
    "hasInternship": false,
    "id": "localadmin",
    "lastName": "Smith",
    "location": "Champaign",
    "major": "Computer Science",
    "programmingAbility": 4,
    "programmingYears": 0,
    "race": null,
    "resumeFilename": "",
    "school": "University of Illinois at Urbana-Champaign",
    "timezone": "Central",
    "updatedAt": 1635291098
}
```

PUT /registration/attendee/
------------------

Updated the registration for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
    "firstName": "John",
    "lastName": "Smith",
    "email": "john@gmail.com",
    "shirtSize": "M",
    "diet": "NONE",
    "age": 19,
    "graduationYear": 2019,
    "transportation": "NONE",
    "school": "University of Illinois at Urbana-Champaign",
    "major": "Computer Science",
    "gender": "MALE",
    "professionalInterest": "INTERNSHIP",
    "github": "JSmith",
    "linkedin": "john-smith",
    "location": "Champaign",
    "interests": "Software",
    "isNovice": false,
    "isPrivate": false,
    "phoneNumber": "555-555-5555",
    "programmingAbility": 4,
    "timezone": "Central",
    "longforms": [
        {
            "response": "This is a longform."
        }
    ],
    "extraInfos": [
        {
            "response": "This is an extra info."
        }
    ],
    "osContributors": [
        {
            "name": "Tom",
            "contactInfo": "tom@gmail.com"
        }
    ],
    "collaborators": [
        {
            "github": "collabgithub"
        }
    ]
}
```

Response format:
```
{
    "createdAt": 1635291098,
    "degreePursued": "",
    "email": "john@gmail.com",
    "firstName": "John",
    "gender": "MALE",
    "github": "localadmin",
    "graduationYear": 2019,
    "hasInternship": false,
    "id": "localadmin",
    "lastName": "Smith",
    "location": "Champaign",
    "major": "Computer Science",
    "programmingAbility": 4,
    "programmingYears": 0,
    "race": null,
    "resumeFilename": "",
    "school": "University of Illinois at Urbana-Champaign",
    "timezone": "Central",
    "updatedAt": 1635293042
}
```

GET /registration/mentor/USERID/
-------------------------

Returns the mentor registration stored for the mentor with the `id` `USERID`.

Response format:
```
{
    "createdAt": 1635290586,
    "email": "john@gmail.com",
    "firstName": "John",
    "github": "localadmin",
    "id": "localadmin",
    "lastName": "Smith",
    "linkedin": "john-smith",
    "shirtSize": "M",
    "updatedAt": 1635290586
}
```

GET /registration/mentor/
-------------------------

Returns the mentor registration stored for the mentor with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
    "createdAt": 1635290586,
    "email": "john@gmail.com",
    "firstName": "John",
    "github": "localadmin",
    "id": "localadmin",
    "lastName": "Smith",
    "linkedin": "john-smith",
    "shirtSize": "M",
    "updatedAt": 1635290586
}
```

POST /registration/mentor/
--------------------------

Creates a registration for the mentor with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
    "firstName": "John",
    "lastName": "Smith",
    "email": "john@gmail.com",
    "shirtSize": "M",
    "github": "JSmith",
    "linkedin": "john-smith"
}
```

Response format:
```
{
    "createdAt": 1635293185,
    "email": "john@gmail.com",
    "firstName": "John",
    "github": "localadmin",
    "id": "localadmin",
    "lastName": "Smith",
    "linkedin": "john-smith",
    "shirtSize": "M",
    "updatedAt": 1635293185
}
```

PUT /registration/mentor/
-------------------------

Updated the registration for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
    "firstName": "John",
    "lastName": "Smith",
    "email": "john@gmail.com",
    "shirtSize": "M",
    "github": "JSmith",
    "linkedin": "john-smith"
}
```

Response format:
```
{
    "createdAt": 1635293185,
    "email": "john@gmail.com",
    "firstName": "John",
    "github": "localadmin",
    "id": "localadmin",
    "lastName": "Smith",
    "linkedin": "john-smith",
    "shirtSize": "M",
    "updatedAt": 1635293241
}
```

GET /registration/attendee/list/?key=value
-----------------------------------

Returns the user registrations, filtered with the given key-value pairs (optional)

Response format:
```
{
    "registrations": [
        {
            "createdAt": 1635291098,
            "degreePursued": "",
            "email": "john@gmail.com",
            "firstName": "John",
            "gender": "MALE",
            "github": "localadmin",
            "graduationYear": 2019,
            "hasInternship": false,
            "id": "localadmin",
            "lastName": "Smith",
            "location": "Champaign",
            "major": "Computer Science",
            "programmingAbility": 4,
            "programmingYears": 0,
            "race": null,
            "resumeFilename": "",
            "school": "University of Illinois at Urbana-Champaign",
            "timezone": "Central",
            "updatedAt": 1635293042
        },
        {
            "createdAt": 1635294123,
            "degreePursued": "",
            "email": "john@gmail.com",
            "firstName": "John",
            "gender": "MALE",
            "github": "test",
            "graduationYear": 2019,
            "hasInternship": false,
            "id": "github000001",
            "lastName": "Smith",
            "location": "Champaign",
            "major": "Computer Science",
            "programmingAbility": 4,
            "programmingYears": 0,
            "race": null,
            "resumeFilename": "",
            "school": "University of Illinois at Urbana-Champaign",
            "timezone": "Central",
            "updatedAt": 1635294123
        }
    ]
}
```

GET /registration/mentor/list/?key=value
-----------------------------------

Returns the mentor registrations, filtered with the given key-value pairs (optional)

Response format:
```
{
    "registrations": [
        {
            "createdAt": 1635293185,
            "email": "john@gmail.com",
            "firstName": "John",
            "github": "localadmin",
            "id": "localadmin",
            "lastName": "Smith",
            "linkedin": "john-smith",
            "shirtSize": "M",
            "updatedAt": 1635293241
        },
        {
            "createdAt": 1635293726,
            "email": "john-doe@gmail.com",
            "firstName": "John",
            "github": "test",
            "id": "github000001",
            "lastName": "Doe",
            "linkedin": "john-doe",
            "shirtSize": "M",
            "updatedAt": 1635293726
        }
    ]
}
```
