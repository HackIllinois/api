Registration
============

!!! warning
	The exact fields in the registration requests and responses **will change** based on the registration definitions provided in the API configuration file.
	Please consult them accordingly.

	`config/production_config.json` contains the most up-to-date version of `REGISTRATION_DEFINITION`,
	`MENTOR_REGISTRATION_DEFINITION`, and `REGISTRATION_STAT_FIELDS`.

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
		"gender": "MALE",
		"email": "john@gmail.com",
		"race": "WHITE",
		"selfTransport": "YES",
		"chicagoPurdueTransport": "N/A",
		"location": "Champaign, IL",
		"school": "University of Illinois at Urbana-Champaign",
		"major": "Computer Science",
		"degreePursued": "BACHELORS",
		"graduationYear": 2025,
		"resumeFilename": "smith-resume.pdf",
		"whyHack": "I want to learn and program. Hack yeah!",
		"programmingYears": 2,
		"programmingAbility": 5,
		"interests": [
			"Company Q&As and Networking events",
			"Meeting new people"
		],
		"outreachSurvey": [
			"CS Department Email"
		],
		"dietary": [
			"Lactose-Intolerant"
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
		"gender": "MALE",
		"email": "john@gmail.com",
		"race": "WHITE",
		"selfTransport": "YES",
		"chicagoPurdueTransport": "N/A",
		"location": "Champaign, IL",
		"school": "University of Illinois at Urbana-Champaign",
		"major": "Computer Science",
		"degreePursued": "BACHELORS",
		"graduationYear": 2025,
		"resumeFilename": "smith-resume.pdf",
		"whyHack": "I want to learn and program. Hack yeah!",
		"programmingYears": 2,
		"programmingAbility": 5,
		"interests": [
			"Company Q&As and Networking events",
			"Meeting new people"
		],
		"outreachSurvey": [
			"CS Department Email"
		],
		"dietary": [
			"Lactose-Intolerant"
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
	"gender": "MALE",
	"email": "john@gmail.com",
	"race": "WHITE",
	"selfTransport": "YES",
	"chicagoPurdueTransport": "N/A",
	"location": "Champaign, IL",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"degreePursued": "BACHELORS",
	"graduationYear": 2025,
	"resumeFilename": "smith-resume.pdf",
	"whyHack": "I want to learn and program. Hack yeah!",
	"programmingYears": 2,
	"programmingAbility": 5,
	"interests": [
		"Company Q&As and Networking events",
		"Meeting new people"
	],
	"outreachSurvey": [
		"CS Department Email"
	],
	"dietary": [
		"Lactose-Intolerant"
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
	"gender": "MALE",
	"email": "john@gmail.com",
	"race": "WHITE",
	"selfTransport": "YES",
	"chicagoPurdueTransport": "N/A",
	"location": "Champaign, IL",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"degreePursued": "BACHELORS",
	"graduationYear": 2025,
	"resumeFilename": "smith-resume.pdf",
	"whyHack": "I want to learn and program. Hack yeah!",
	"programmingYears": 2,
	"programmingAbility": 5,
	"interests": [
		"Company Q&As and Networking events",
		"Meeting new people"
	],
	"outreachSurvey": [
		"CS Department Email"
	],
	"dietary": [
		"Lactose-Intolerant"
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
	"gender": "MALE",
	"email": "john@gmail.com",
	"race": "WHITE",
	"selfTransport": "YES",
	"chicagoPurdueTransport": "N/A",
	"location": "Champaign, IL",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"degreePursued": "BACHELORS",
	"graduationYear": 2025,
	"resumeFilename": "smith-resume.pdf",
	"whyHack": "I want to learn and program. Hack yeah!",
	"programmingYears": 2,
	"programmingAbility": 5,
	"interests": [
		"Company Q&As and Networking events",
		"Meeting new people"
	],
	"outreachSurvey": [
		"CS Department Email"
	],
	"dietary": [
		"Lactose-Intolerant"
	]
}
```

```json title="Example response"
{
	"id": "github0000001",
	"firstName": "John",
	"lastName": "Smith",
	"gender": "MALE",
	"email": "john@gmail.com",
	"race": "WHITE",
	"selfTransport": "YES",
	"chicagoPurdueTransport": "N/A",
	"location": "Champaign, IL",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"degreePursued": "BACHELORS",
	"graduationYear": 2025,
	"resumeFilename": "smith-resume.pdf",
	"whyHack": "I want to learn and program. Hack yeah!",
	"programmingYears": 2,
	"programmingAbility": 5,
	"interests": [
		"Company Q&As and Networking events",
		"Meeting new people"
	],
	"outreachSurvey": [
		"CS Department Email"
	],
	"dietary": [
		"Lactose-Intolerant"
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
	"gender": "MALE",
	"email": "john@gmail.com",
	"race": "WHITE",
	"selfTransport": "YES",
	"chicagoPurdueTransport": "N/A",
	"location": "Champaign, IL",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"degreePursued": "BACHELORS",
	"graduationYear": 2025,
	"resumeFilename": "smith-resume.pdf",
	"whyHack": "I want to learn and program. Hack yeah!",
	"programmingYears": 2,
	"programmingAbility": 5,
	"interests": [
		"Company Q&As and Networking events",
		"Meeting new people"
	],
	"outreachSurvey": [
		"CS Department Email"
	],
	"dietary": [
		"Lactose-Intolerant"
	]
}
```

```json title="Example response"
{
	"id": "github0000001",
	"firstName": "John",
	"lastName": "Smith",
	"gender": "MALE",
	"email": "john@gmail.com",
	"race": "WHITE",
	"selfTransport": "YES",
	"chicagoPurdueTransport": "N/A",
	"location": "Champaign, IL",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"degreePursued": "BACHELORS",
	"graduationYear": 2025,
	"resumeFilename": "smith-resume.pdf",
	"whyHack": "I want to learn and program. Hack yeah!",
	"programmingYears": 2,
	"programmingAbility": 5,
	"interests": [
		"Company Q&As and Networking events",
		"Meeting new people"
	],
	"outreachSurvey": [
		"CS Department Email"
	],
	"dietary": [
		"Lactose-Intolerant"
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
			"gender": "MALE",
			"email": "john@gmail.com",
			"race": "WHITE",
			"selfTransport": "YES",
			"chicagoPurdueTransport": "N/A",
			"location": "Champaign, IL",
			"school": "University of Illinois at Urbana-Champaign",
			"major": "Computer Science",
			"degreePursued": "BACHELORS",
			"graduationYear": 2025,
			"resumeFilename": "smith-resume.pdf",
			"whyHack": "I want to learn and program. Hack yeah!",
			"programmingYears": 2,
			"programmingAbility": 5,
			"interests": [
				"Company Q&As and Networking events",
				"Meeting new people"
			],
			"outreachSurvey": [
				"CS Department Email"
			],
			"dietary": [
				"Lactose-Intolerant"
			]
		},
		{
			"id": "github0000002",
			"firstName": "John",
			"lastName": "Doe",
			"gender": "MALE",
			"email": "jdoe@gmail.com",
			"race": "MULTIRACIAL",
			"selfTransport": "YES",
			"chicagoPurdueTransport": "N/A",
			"location": "Champaign, IL",
			"school": "University of Illinois at Urbana-Champaign",
			"major": "Computer Science",
			"degreePursued": "MASTERS",
			"graduationYear": 2023,
			"resumeFilename": "doe-resume.pdf",
			"whyHack": "I also want to learn and program. Hack yeah!",
			"programmingYears": 6,
			"programmingAbility": 8,
			"interests": [
				"Company Q&As and Networking events"
			],
			"outreachSurvey": [
				"CS Department Email"
			],
			"dietary": []
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
