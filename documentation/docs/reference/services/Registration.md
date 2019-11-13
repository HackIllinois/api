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
		"major": "Computer Science",
		"degreePursued": "Bachelor’s",
		"graduationYear": 2020,
		"careerInterest": ["INTERNSHIP"],
		"resumeFilename": "resume.pdf",
		"programmingYears": 10,
		"programmingAbility": 7,
		"isOSContributor": true,
		"categoryInterests": ["Systems", "Web"],
		"languageInterests": ["JavaScript", "Python"],
		"needsBus": false,
		"hasAttended": true,
		"howDiscovered": ["Peer"],
		"shirtSize": "L",
		"dietaryRestrictions": ["NONE"],
		"hasDisability": false
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
		"id": "github0000001",
		"email": "john@gmail.com",
		"github": "jsmith",
		"createdAt": 0000000001,
		"updatedAt": 0000000002,
		"firstName": "John",
		"lastName": "Smith",
		"gender": "MALE",
		"school": "University of Illinois at Urbana-Champaign",
		"major": "Computer Science",
		"degreePursued": "Bachelor’s",
		"graduationYear": 2020,
		"careerInterest": ["INTERNSHIP"],
		"resumeFilename": "resume.pdf",
		"programmingYears": 10,
		"programmingAbility": 7,
		"isOSContributor": true,
		"categoryInterests": ["Systems", "Web"],
		"languageInterests": ["JavaScript", "Python"],
		"needsBus": false,
		"hasAttended": true,
		"howDiscovered": ["Peer"],
		"shirtSize": "L",
		"dietaryRestrictions": ["NONE"],
		"hasDisability": false
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
	"id": "github0000001",
	"email": "john@gmail.com",
	"github": "jsmith",
	"createdAt": 0000000001,
	"updatedAt": 0000000002,
	"firstName": "John",
	"lastName": "Smith",
	"gender": "MALE",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"degreePursued": "Bachelor’s",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",
	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["Systems", "Web"],
	"languageInterests": ["JavaScript", "Python"],
	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["Peer"],
	"shirtSize": "L",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
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
	"major": "Computer Science",
	"degreePursued": "Bachelor’s",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",
	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["Systems", "Web"],
	"languageInterests": ["JavaScript", "Python"],
	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["Peer"],
	"shirtSize": "L",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
}
```

POST /registration/attendee/
-------------------

Creates a registration for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
	"email": "john@gmail.com",
	"github": "jsmith",
	"firstName": "John",
	"lastName": "Smith",
	"gender": "MALE",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"degreePursued": "Bachelor’s",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",
	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["Systems", "Web"],
	"languageInterests": ["JavaScript", "Python"],
	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["Peer"],
	"shirtSize": "L",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
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
	"major": "Computer Science",
	"degreePursued": "Bachelor’s",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",
	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["Systems", "Web"],
	"languageInterests": ["JavaScript", "Python"],
	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["Peer"],
	"shirtSize": "L",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
}
```

PUT /registration/attendee/
------------------

Updated the registration for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
	"email": "john@gmail.com",
	"github": "jsmith",
	"firstName": "John",
	"lastName": "Smith",
	"gender": "MALE",
	"school": "University of Illinois at Urbana-Champaign",
	"major": "Computer Science",
	"degreePursued": "Bachelor’s",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",
	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["Systems", "Web"],
	"languageInterests": ["JavaScript", "Python"],
	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["Peer"],
	"shirtSize": "L",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
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
	"major": "Computer Science",
	"degreePursued": "Bachelor’s",
	"graduationYear": 2020,
	"careerInterest": ["INTERNSHIP"],
	"resumeFilename": "resume.pdf",
	"programmingYears": 10,
	"programmingAbility": 7,
	"isOSContributor": true,
	"categoryInterests": ["Systems", "Web"],
	"languageInterests": ["JavaScript", "Python"],
	"needsBus": false,
	"hasAttended": true,
	"howDiscovered": ["Peer"],
	"shirtSize": "L",
	"dietaryRestrictions": ["NONE"],
	"hasDisability": false
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
			"id": "github0000001",
			"email": "john@gmail.com",
			"github": "jsmith",
			"createdAt": 0000000001,
			"updatedAt": 0000000002,
			"firstName": "John",
			"lastName": "Smith",
			"gender": "MALE",
			"school": "University of Illinois at Urbana-Champaign",
			"major": "Computer Science",
			"degreePursued": "Bachelor’s",
			"graduationYear": 2020,
			"careerInterest": ["INTERNSHIP"],
			"resumeFilename": "resume.pdf",
			"programmingYears": 10,
			"programmingAbility": 7,
			"isOSContributor": true,
			"categoryInterests": ["Systems", "Web"],
			"languageInterests": ["JavaScript", "Python"],
			"needsBus": false,
			"hasAttended": true,
			"howDiscovered": ["Peer"],
			"shirtSize": "L",
			"dietaryRestrictions": ["NONE"],
			"hasDisability": false
		},
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
			"major": "Computer Science",
			"degreePursued": "Bachelor’s",
			"graduationYear": 2020,
			"careerInterest": ["INTERNSHIP"],
			"resumeFilename": "resume.pdf",
			"programmingYears": 10,
			"programmingAbility": 7,
			"isOSContributor": true,
			"categoryInterests": ["Systems", "Web"],
			"languageInterests": ["JavaScript", "Python"],
			"needsBus": false,
			"hasAttended": true,
			"howDiscovered": ["Peer"],
			"shirtSize": "L",
			"dietaryRestrictions": ["NONE"],
			"hasDisability": false
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
