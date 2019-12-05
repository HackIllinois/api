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
		"id": "github0000001",
		"email": "john@gmail.com",
		"github": "jsmith",
		"createdAt": 0000000001,
		"updatedAt": 0000000002,

		"firstName": "John",
		"lastName": "Smith",
		"gender": "MALE",

		"school": "University of Illinois at Urbana-Champaign",
		"major": "CS",
		"degreePursued": "BACHELORS",
		"graduationYear": 2020,
		"careerInterest": ["INTERNSHIP"],
		"resumeFilename": "resume.pdf",

		"programmingYears": 10,
		"programmingAbility": 7,
		"isOSContributor": true,
		"categoryInterests": ["SYSTEMS", "WEBDEV"],
		"languageInterests": ["JavaScript", "Python"],

		"needsBus": false,
		"hasAttended": true,
		"howDiscovered": ["FRIEND"]
	},
	"mentor": {
		"id": "github0000001",
		"email": "john@gmail.com",
		"github": "JSmith",
		"createdAt": 0000000001,
		"updatedAt": 0000000002,

		"firstName": "John",
		"lastName": "Smith",
		"photoFilename": "profile.jpg",
		"biography": "I write code.",
		"projectName": "HackIllinois",
		"projectDescription": "The best hackathon ever!",

		"shirtSize": "M",
		"dietaryRestrictions": ["NONE"],
		"hasDisability": false
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
		"id": "github0000001",
		"email": "john@gmail.com",
		"github": "jsmith",
		"createdAt": 0000000001,
		"updatedAt": 0000000002,

		"firstName": "John",
		"lastName": "Smith",
		"gender": "MALE",

		"school": "University of Illinois at Urbana-Champaign",
		"major": "CS",
		"degreePursued": "BACHELORS",
		"graduationYear": 2020,
		"careerInterest": ["INTERNSHIP"],
		"resumeFilename": "resume.pdf",

		"programmingYears": 10,
		"programmingAbility": 7,
		"isOSContributor": true,
		"categoryInterests": ["SYSTEMS", "WEBDEV"],
		"languageInterests": ["JavaScript", "Python"],

		"needsBus": false,
		"hasAttended": true,
		"howDiscovered": ["FRIEND"]
	},
	"mentor": {
		"id": "github0000001",
		"email": "john@gmail.com",
		"github": "JSmith",
		"createdAt": 0000000001,
		"updatedAt": 0000000002,

		"firstName": "John",
		"lastName": "Smith",
		"photoFilename": "profile.jpg",
		"biography": "I write code.",
		"projectName": "HackIllinois",
		"projectDescription": "The best hackathon ever!",

		"shirtSize": "M",
		"dietaryRestrictions": ["NONE"],
		"hasDisability": false
	}
}
```

GET /registration/attendee/USERID/
-------------------------

Returns the user registration stored for the user with the `id` `USERID`.

Response format:
```
{
	"id": "github0000001",
	"email": "john@gmail.com",
	"github": "jsmith",
	"createdAt": 0000000001,
	"updatedAt": 0000000002,

	"firstName": "John",
	"lastName": "Smith",
	"gender": "MALE",

	"school": "University of Illinois at Urbana-Champaign",
	"major": "CS",
	"degreePursued": "BACHELORS",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",

	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["SYSTEMS", "WEBDEV"],
	"languageInterests": ["JavaScript", "Python"],

	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["FRIEND"]
}
```

GET /registration/attendee/
------------------

Returns the user registration stored for the user with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
	"id": "github0000001",
	"email": "john@gmail.com",
	"github": "jsmith",
	"createdAt": 0000000001,
	"updatedAt": 0000000002,

	"firstName": "John",
	"lastName": "Smith",
	"gender": "MALE",

	"school": "University of Illinois at Urbana-Champaign",
	"major": "CS",
	"degreePursued": "BACHELORS",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",

	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["SYSTEMS", "WEBDEV"],
	"languageInterests": ["JavaScript", "Python"],

	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["FRIEND"]
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
	"gender": "MALE",

	"school": "University of Illinois at Urbana-Champaign",
	"major": "CS",
	"degreePursued": "BACHELORS",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",

	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["SYSTEMS", "WEBDEV"],
	"languageInterests": ["JavaScript", "Python"],

	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["FRIEND"]
}
```

Response format:
```
{
	"id": "github0000001",
	"email": "john@gmail.com",
	"github": "jsmith",
	"createdAt": 0000000001,
	"updatedAt": 0000000002,

	"firstName": "John",
	"lastName": "Smith",
	"gender": "MALE",

	"school": "University of Illinois at Urbana-Champaign",
	"major": "CS",
	"degreePursued": "BACHELORS",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",

	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["SYSTEMS", "WEBDEV"],
	"languageInterests": ["JavaScript", "Python"],

	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["FRIEND"]
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
	"gender": "MALE",

	"school": "University of Illinois at Urbana-Champaign",
	"major": "CS",
	"degreePursued": "BACHELORS",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",

	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["SYSTEMS", "WEBDEV"],
	"languageInterests": ["JavaScript", "Python"],

	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["FRIEND"]
}
```

Response format:
```
{
	"id": "github0000001",
	"email": "john@gmail.com",
	"github": "jsmith",
	"createdAt": 0000000001,
	"updatedAt": 0000000002,

	"firstName": "John",
	"lastName": "Smith",
	"gender": "MALE",

	"school": "University of Illinois at Urbana-Champaign",
	"major": "CS",
	"degreePursued": "BACHELORS",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",

	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["SYSTEMS", "WEBDEV"],
	"languageInterests": ["JavaScript", "Python"],

	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["FRIEND"]
}
```

GET /registration/mentor/USERID/
-------------------------

Returns the mentor registration stored for the mentor with the `id` `USERID`.

Response format:
```
{
	"id": "github0000001",
	"email": "john@gmail.com",
	"github": "JSmith",
	"createdAt": 0000000001,
	"updatedAt": 0000000002,

	"firstName": "John",
	"lastName": "Smith",
	"photoFilename": "profile.jpg",
	"biography": "I write code.",
	"projectName": "HackIllinois",
	"projectDescription": "The best hackathon ever!",

	"shirtSize": "M",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
}
```

GET /registration/mentor/
-------------------------

Returns the mentor registration stored for the mentor with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
	"id": "github0000001",
	"email": "john@gmail.com",
	"github": "JSmith",
	"createdAt": 0000000001,
	"updatedAt": 0000000002,

	"firstName": "John",
	"lastName": "Smith",
	"photoFilename": "profile.jpg",
	"biography": "I write code.",
	"projectName": "HackIllinois",
	"projectDescription": "The best hackathon ever!",

	"shirtSize": "M",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
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
	"photoFilename": "profile.jpg",
	"biography": "I write code.",
	"projectName": "HackIllinois",
	"projectDescription": "The best hackathon ever!",

	"shirtSize": "M",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
}
```

Response format:
```
{
	"id": "github0000001",
	"email": "john@gmail.com",
	"github": "JSmith",
	"createdAt": 0000000001,
	"updatedAt": 0000000002,

	"firstName": "John",
	"lastName": "Smith",
	"photoFilename": "profile.jpg",
	"biography": "I write code.",
	"projectName": "HackIllinois",
	"projectDescription": "The best hackathon ever!",

	"shirtSize": "M",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
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
	"photoFilename": "profile.jpg",
	"biography": "I write code.",
	"projectName": "HackIllinois",
	"projectDescription": "The best hackathon ever!",

	"shirtSize": "M",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
}
```

Response format:
```
{
	"id": "github0000001",
	"email": "john@gmail.com",
	"github": "JSmith",
	"createdAt": 0000000001,
	"updatedAt": 0000000002,

	"firstName": "John",
	"lastName": "Smith",
	"photoFilename": "profile.jpg",
	"biography": "I write code.",
	"projectName": "HackIllinois",
	"projectDescription": "The best hackathon ever!",

	"shirtSize": "M",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
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
			"id": "github0000001",
			"email": "john@gmail.com",
			"github": "jsmith",
			"createdAt": 0000000001,
			"updatedAt": 0000000002,

			"firstName": "John",
			"lastName": "Smith",
			"gender": "MALE",

			"school": "University of Illinois at Urbana-Champaign",
			"major": "CS",
			"degreePursued": "BACHELORS",
			"graduationYear": 2020,
			"careerInterest": ["INTERNSHIP"],
			"resumeFilename": "resume.pdf",

			"programmingYears": 10,
			"programmingAbility": 7,
			"isOSContributor": true,
			"categoryInterests": ["SYSTEMS", "WEBDEV"],
			"languageInterests": ["JavaScript", "Python"],

			"needsBus": false,
			"hasAttended": true,
			"howDiscovered": ["FRIEND"]
		},
		{
			"id": "github0000002",
			"email": "john2@gmail.com",
			"github": "jsmith2",
			"createdAt": 0000000001,
			"updatedAt": 0000000002,

			"firstName": "John2",
			"lastName": "Smith2",
			"gender": "MALE",

			"school": "University of Illinois at Urbana-Champaign",
			"major": "CS",
			"degreePursued": "BACHELORS",
			"graduationYear": 2020,
			"careerInterest": ["INTERNSHIP"],
			"resumeFilename": "resume.pdf",

			"programmingYears": 10,
			"programmingAbility": 7,
			"isOSContributor": true,
			"categoryInterests": ["SYSTEMS", "WEBDEV"],
			"languageInterests": ["JavaScript", "Python"],

			"needsBus": false,
			"hasAttended": true,
			"howDiscovered": ["FRIEND"]
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
			"id": "github0000001",
			"email": "john@gmail.com",
			"github": "JSmith",
			"createdAt": 0000000001,
			"updatedAt": 0000000002,

			"firstName": "John",
			"lastName": "Smith",
			"photoFilename": "profile.jpg",
			"biography": "I write code.",
			"projectName": "HackIllinois",
			"projectDescription": "The best hackathon ever!",

			"shirtSize": "M",
			"dietaryRestrictions": ["NONE"],
			"hasDisability": false
		},
		{
			"id": "github0000002",
			"email": "john2@gmail.com",
			"github": "JSmith2",
			"createdAt": 0000000001,
			"updatedAt": 0000000002,

			"firstName": "John2",
			"lastName": "Smith2",
			"photoFilename": "profile2.jpg",
			"biography": "I write code.",
			"projectName": "HackIllinois",
			"projectDescription": "The best hackathon ever!",

			"shirtSize": "M",
			"dietaryRestrictions": ["NONE"],
			"hasDisability": false
		}
	]
}
```
