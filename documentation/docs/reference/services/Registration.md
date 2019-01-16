Registration
============

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
		"interests": "INTERNSHIP",
		"github": "JSmith",
		"linkedin": "john-smith",
		"interests": "Software",
		"isBeginner": false,
		"priorAttendance": false,
		"phone": "555-555-5555",
		"extraInfo": "Extra information",
		"teamMembers": [
			"member1",
			"member2",
			"member3
		],
		"createdAt": 123123,
		"updatedAt": 1234353,
		"beginnerInfo": {
				"versionControl": 4,
				"pullRequest": 2,
				"yearsExperience": 6,
				"technicalSkills": [
					"algorithms",
					"distributed systems",
					"machine learning"
				]
		},
		{
			"isOSContributor": true
		}
	},
	"mentor": {
		"id": "github0001",
		"firstName": "John",
		"lastName": "Smith",
		"email": "john.smith@email.com"
		"shirtSize": "S",
		"github": "JohnSmith",
		"linkedin": "john-smith",
		"createdAt": 1231231,
		"updatedAt": 3132423
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
		"interests": "INTERNSHIP",
		"github": "JSmith",
		"linkedin": "john-smith",
		"interests": "Software",
		"isBeginner": false,
		"priorAttendance": false,
		"phone": "555-555-5555",
		"extraInfo": "Extra information",
		"teamMembers": [
			"member1",
			"member2",
			"member3
		],
		"createdAt": 123123,
		"updatedAt": 1234353,
		"beginnerInfo": {
				"versionControl": 4,
				"pullRequest": 2,
				"yearsExperience": 6,
				"technicalSkills": [
					"algorithms",
					"distributed systems",
					"machine learning"
				]
		},
		"isOSContributor": true
	},
	"mentor": {
		"id": "github0001",
		"firstName": "John",
		"lastName": "Smith",
		"email": "john.smith@email.com"
		"shirtSize": "S",
		"github": "JohnSmith",
		"linkedin": "john-smith",
		"createdAt": 1231231,
		"updatedAt": 3132423
	}
}
```

GET /registration/attendee/USERID/
-------------------------

Returns the user registration stored for the Attendee with the `id` `USERID`.

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
	"interests": "INTERNSHIP",
	"github": "JSmith",
	"linkedin": "john-smith",
	"interests": "Software",
	"isBeginner": false,
	"priorAttendance": false,
	"phone": "555-555-5555",
	"extraInfo": "Extra information",
	"teamMembers": [
		"member1",
		"member2",
		"member3
	],
	"createdAt": 123123,
	"updatedAt": 1234353,
	"beginnerInfo": {
			"versionControl": 4,
			"pullRequest": 2,
			"yearsExperience": 6,
			"technicalSkills": [
				"algorithms",
				"distributed systems",
				"machine learning"
			]
	},
	"isOSContributor": true
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
	"interests": "INTERNSHIP",
	"github": "JSmith",
	"linkedin": "john-smith",
	"interests": "Software",
	"isBeginner": false,
	"priorAttendance": false,
	"phone": "555-555-5555",
	"extraInfo": "Extra information",
	"teamMembers": [
		"member1",
		"member2",
		"member3
	],
	"createdAt": 123123,
	"updatedAt": 1234353,
	"beginnerInfo": {
			"versionControl": 4,
			"pullRequest": 2,
			"yearsExperience": 6,
			"technicalSkills": [
				"algorithms",
				"distributed systems",
				"machine learning"
			]
	},
	"isOSContributor": true
}
```

POST /registration/attendee/
-------------------

Creates a registration for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
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
	"interests": "INTERNSHIP",
	"github": "JSmith",
	"linkedin": "john-smith",
	"interests": "Software",
	"isBeginner": false,
	"priorAttendance": false,
	"phone": "555-555-5555",
	"extraInfo": "Extra information",
	"teamMembers": [
		"member1",
		"member2",
		"member3
	],
	"createdAt": 123123,
	"updatedAt": 1234353,
	"beginnerInfo": {
			"versionControl": 4,
			"pullRequest": 2,
			"yearsExperience": 6,
			"technicalSkills": [
				"algorithms",
				"distributed systems",
				"machine learning"
			]
	},
	"isOSContributor": true
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
	"interests": "INTERNSHIP",
	"github": "JSmith",
	"linkedin": "john-smith",
	"interests": "Software",
	"isBeginner": false,
	"priorAttendance": false,
	"phone": "555-555-5555",
	"extraInfo": "Extra information",
	"teamMembers": [
		"member1",
		"member2",
		"member3
	],
	"createdAt": 123123,
	"updatedAt": 1234353,
	"beginnerInfo": {
			"versionControl": 4,
			"pullRequest": 2,
			"yearsExperience": 6,
			"technicalSkills": [
				"algorithms",
				"distributed systems",
				"machine learning"
			]
	},
	"isOSContributor": true
}
```

PUT /registration/attendee/
------------------

Updated the registration for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
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
	"interests": "INTERNSHIP",
	"github": "JSmith",
	"linkedin": "john-smith",
	"interests": "Software",
	"isBeginner": false,
	"priorAttendance": false,
	"phone": "555-555-5555",
	"extraInfo": "Extra information",
	"teamMembers": [
		"member1",
		"member2",
		"member3
	],
	"createdAt": 123123,
	"updatedAt": 1234353,
	"beginnerInfo": {
			"versionControl": 4,
			"pullRequest": 2,
			"yearsExperience": 6,
			"technicalSkills": [
				"algorithms",
				"distributed systems",
				"machine learning"
			]
	},
	"isOSContributor": true
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
	"interests": "INTERNSHIP",
	"github": "JSmith",
	"linkedin": "john-smith",
	"interests": "Software",
	"isBeginner": false,
	"priorAttendance": false,
	"phone": "555-555-5555",
	"extraInfo": "Extra information",
	"teamMembers": [
		"member1",
		"member2",
		"member3
	],
	"createdAt": 123123,
	"updatedAt": 1234353,
	"beginnerInfo": {
			"versionControl": 4,
			"pullRequest": 2,
			"yearsExperience": 6,
			"technicalSkills": [
				"algorithms",
				"distributed systems",
				"machine learning"
			]
	},
	"isOSContributor": true
}
```

GET /registration/mentor/USERID/
-------------------------

Returns the mentor registration stored for the mentor with the `id` `USERID`.

Response format:
```
{
	"id": "github0001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john.smith@email.com"
	"shirtSize": "S",
	"github": "JohnSmith",
	"linkedin": "john-smith",
	"createdAt": 1231231,
	"updatedAt": 3132423
}
```

GET /registration/mentor/
-------------------------

Returns the mentor registration stored for the mentor with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
	"id": "github0001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john.smith@email.com"
	"shirtSize": "S",
	"github": "JohnSmith",
	"linkedin": "john-smith",
	"createdAt": 1231231,
	"updatedAt": 3132423
}
```

POST /registration/mentor/
--------------------------

Creates a registration for the mentor with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
	"id": "github0001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john.smith@email.com"
	"shirtSize": "S",
	"github": "JohnSmith",
	"linkedin": "john-smith",
	"createdAt": 1231231,
	"updatedAt": 3132423
}
```

Response format:
```
{
	"id": "github0001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john.smith@email.com"
	"shirtSize": "S",
	"github": "JohnSmith",
	"linkedin": "john-smith",
	"createdAt": 1231231,
	"updatedAt": 3132423
}
```

PUT /registration/mentor/
-------------------------

Updated the registration for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
	"id": "github0001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john.smith@email.com"
	"shirtSize": "S",
	"github": "JohnSmith",
	"linkedin": "john-smith",
	"createdAt": 1231231,
	"updatedAt": 3132423
}
```

Response format:
```
{
	"id": "github0001",
	"firstName": "John",
	"lastName": "Smith",
	"email": "john.smith@email.com"
	"shirtSize": "S",
	"github": "JohnSmith",
	"linkedin": "john-smith",
	"createdAt": 1231231,
	"updatedAt": 3132423
}
```

GET /registration/filter/?key=value
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
			"interests": "INTERNSHIP",
			"github": "JSmith",
			"linkedin": "john-smith",
			"interests": "Software",
			"isBeginner": false,
			"priorAttendance": false,
			"phone": "555-555-5555",
			"extraInfo": "Extra information",
			"teamMembers": [
				"member1",
				"member2",
				"member3
			],
			"createdAt": 123123,
			"updatedAt": 1234353,
			"beginnerInfo": {
					"versionControl": 4,
					"pullRequest": 2,
					"yearsExperience": 6,
					"technicalSkills": [
						"algorithms",
						"distributed systems",
						"machine learning"
					]
			},
			"isOSContributor": true
		},
		{
			"id": "github0000002"
			"firstName": "John2",
			"lastName": "Smith2",
			"email": "john2@gmail.com",
			"shirtSize": "L",
			"diet": "NONE",
			"age": 19,
			"graduationYear": 2019,
			"transportation": "NONE",
			"school": "University of Illinois at Urbana-Champaign",
			"major": "Computer Science",
			"gender": "MALE",
			"interests": "INTERNSHIP",
			"github": "JSmith",
			"linkedin": "john-smith",
			"interests": "Software",
			"isBeginner": false,
			"priorAttendance": true,
			"phone": "555-555-5555",
			"extraInfo": "Extra information",
			"teamMembers": [
				"member1",
				"member2",
				"member3
			],
			"createdAt": 123123,
			"updatedAt": 1234353,
			"beginnerInfo": {
					"versionControl": 4,
					"pullRequest": 2,
					"yearsExperience": 6,
					"technicalSkills": [
						"algorithms",
						"distributed systems",
						"machine learning"
					]
			},
			"isOSContributor": true
		}
	]
}
```
