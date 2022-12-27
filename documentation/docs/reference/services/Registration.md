Registration
============

!!! warning
	The exact fields in the registration requests and responses **will change** based on the registration definitions provided in the API configuration file.
	Please consult them accordingly.

GET /registration/
-------------------------

Returns all the registrations stored for the currently authenticated user (determined by the JWT in the `Authorization` header).
If registrations are not found for either Attendee or Mentor, that field is set to null.

Request requires no body.

```json title="Example response"
{
	"attendee": {
		"id": "github0000001",
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
		"interests": "Software",
		"isNovice": false,
		"isPrivate": false,
		"phoneNumber": "555-555-5555",
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
	},
	"mentor": {
		"id": "github0000001",
		"firstName": "John",
		"lastName": "Smith",
		"email": "john@gmail.com",
		"shirtSize": "M",
		"github": "JSmith",
		"linkedin": "john-smith"
	} 
}
```

GET /registration/USERID/
-------------------------

Returns all registrations stored for the user with the `id` `USERID`.
If registrations are not found for either Attendee or Mentor, that field is set to null.

Request requires no body.

```json title="Example response"
{
	"attendee": {
		"id": "github0000001",
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
		"interests": "Software",
		"isNovice": false,
		"isPrivate": false,
		"phoneNumber": "555-555-5555",
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
	},
	"mentor": {
		"id": "github0000001",
		"firstName": "John",
		"lastName": "Smith",
		"email": "john@gmail.com",
		"shirtSize": "M",
		"github": "JSmith",
		"linkedin": "john-smith"
	}
}
```

GET /registration/attendee/USERID/
-------------------------

Returns the user registration stored for the user with the `id` `USERID`.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
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
	"interests": "Software",
	"isNovice": false,
	"isPrivate": false,
	"phoneNumber": "555-555-5555",
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

GET /registration/attendee/
------------------

Returns the user registration stored for the currently authenticated user (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
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
	"interests": "Software",
	"isNovice": false,
	"isPrivate": false,
	"phoneNumber": "555-555-5555",
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

POST /registration/attendee/
-------------------

Creates a registration for the currently authenticated user (determined by the JWT in the `Authorization` header).

```json title="Example request"
{
	"firstName": "John",
	"lastName": "Smith",
	"email": "john@gmail.com",
	"github": "JohnSmith",
	"location": "Urbana, IL",
	"timezone": "CST",
	"gender": [
		"Prefer not to answer"
	],
	"race": "MALE",
	"graduationYear": 2019,
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"graduationYear": 2019,
	"programmingYears": 5,
	"programmingAbility": 10,
	"interests": [
		"corporate engagement",
		"learning"
	]
}
```

```json title="Example response"
{
	"id": "github0000001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john@gmail.com",
	"github": "JohnSmith",
	"location": "Urbana, IL",
	"timezone": "CST",
	"gender": [
		"Prefer not to answer"
	],
	"race": "MALE",
	"graduationYear": 2019,
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"graduationYear": 2019,
	"programmingYears": 5,
	"programmingAbility": 10,
	"interests": [
		"corporate engagement",
		"learning"
	]
}
```

PUT /registration/attendee/
------------------

Update the registration for the currently authenticated user (determined by the JWT in the `Authorization` header).

```json title="Example request"
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
	"interests": "Software",
	"isNovice": false,
	"isPrivate": false,
	"phoneNumber": "555-555-5555",
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

```json title="Example response"
{
	"id": "github0000001",
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
	"interests": "Software",
	"isNovice": false,
	"isPrivate": false,
	"phoneNumber": "555-555-5555",
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

GET /registration/mentor/USERID/
-------------------------

Returns the mentor registration stored for the mentor with the `id` `USERID`.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john@gmail.com",
	"shirtSize": "M",
	"github": "JSmith",
	"linkedin": "john-smith"
}
```

GET /registration/mentor/
-------------------------

Returns the mentor registration stored for the currently authenticated mentor (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john@gmail.com",
	"shirtSize": "M",
	"github": "JSmith",
	"linkedin": "john-smith"
}
```

POST /registration/mentor/
--------------------------

Creates a registration for the currently authenticated mentor (determined by the JWT in the `Authorization` header)

```json title="Example request"
{
	"firstName": "John",
	"lastName": "Smith",
	"email": "john@gmail.com",
	"shirtSize": "M",
	"github": "JSmith",
	"linkedin": "john-smith"
}
```

```json title="Example response"
{
	"id": "github0000001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john@gmail.com",
	"shirtSize": "M",
	"github": "JSmith",
	"linkedin": "john-smith"
}
```

PUT /registration/mentor/
-------------------------

Updated the registration for the currently authenticated user (determined by the JWT in the `Authorization` header).

```json title="Example request"
{
	"firstName": "John",
	"lastName": "Smith",
	"email": "john@gmail.com",
	"shirtSize": "M",
	"github": "JSmith",
	"linkedin": "john-smith"
}
```

```json title="Example response"
{
	"id": "github0000001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john@gmail.com",
	"shirtSize": "M",
	"github": "JSmith",
	"linkedin": "john-smith"
}
```

GET /registration/attendee/list/?key=value
-----------------------------------

Returns the user registrations, filtered with the given key-value pairs (optional)

Request requires no body.

```json title="Example response"
{
	"registrations": [
		{
			"id": "github0000001",
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
			"interests": "Software",
			"isNovice": false,
			"isPrivate": false,
			"phoneNumber": "555-555-5555",
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
		},
		{
			"id": "github0000002",
			"firstName": "John",
			"lastName": "Doe",
			"email": "jdoe@gmail.com",
			"shirtSize": "M",
			"diet": "NONE",
			"age": 19,
			"graduationYear": 2019,
			"transportation": "NONE",
			"school": "Purdue",
			"major": "Computer Science",
			"gender": "MALE",
			"professionalInterest": "INTERNSHIP",
			"github": "JDoe",
			"linkedin": "john-doe",
			"interests": "Software",
			"isNovice": false,
			"isPrivate": false,
			"phoneNumber": "666-666-6666",
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
	]
}
```

GET /registration/mentor/list/?key=value
-----------------------------------

Returns the mentor registrations, filtered with the given key-value pairs (optional)

Request requires no body.

```json title="Example response"
{
	"registrations": [
		{
			"id": "github0000001",
			"firstName": "John",
			"lastName": "Smith",
			"email": "john@gmail.com",
			"shirtSize": "M",
			"github": "JSmith",
			"linkedin": "john-smith"
		},
		{
			"id": "github0000002",
			"firstName": "John",
			"lastName": "Doe",
			"email": "jdoe@gmail.com",
			"shirtSize": "M",
			"github": "JDoe",
			"linkedin": "john-doe"
		}
	]
}
```
