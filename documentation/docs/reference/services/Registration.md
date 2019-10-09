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
		"id": "github0000001"
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
		"id": "github0000001"
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

Response format:
```
{
	"attendee": {
		"id": "github0000001"
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
		"id": "github0000001"
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

Response format:
```
{
	"id": "github0000001"
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

Returns the user registration stored for the user with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
	"id": "github0000001"
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

Response format:
```
{
	"id": "github0000001"
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

Response format:
```
{
	"id": "github0000001"
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

Response format:
```
{
	"id": "github0000001"
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

Returns the mentor registration stored for the mentor with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
	"id": "github0000001"
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
	"id": "github0000001"
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
	"id": "github0000001"
	"firstName": "John",
	"lastName": "Smith",
	"email": "john@gmail.com",
	"shirtSize": "M",
	"github": "JSmith",
	"linkedin": "john-smith"
}
```

GET /registration/attendee/filter/?key=value
-----------------------------------

Returns the user registrations, filtered with the given key-value pairs

Response format:
```
{
	"registrations": [
		{
			"id": "github0000001"
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
			"id": "github0000002"
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

GET /registration/mentor/filter/?key=value
-----------------------------------

Returns the mentor registrations, filtered with the given key-value pairs

Response format:
```
{
	"registrations": [
		{
			"id": "github0000001"
			"firstName": "John",
			"lastName": "Smith",
			"email": "john@gmail.com",
			"shirtSize": "M",
			"github": "JSmith",
			"linkedin": "john-smith"
		},
		{
			"id": "github0000002"
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
