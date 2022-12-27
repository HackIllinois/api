Project
=====

GET /project/PROJECTID/
---------------------

Returns the project with the id of `PROJECTID`.

Request requires no body.

```json title="Example response"
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"description": "Example Project Description",
	"mentors": ["Jane Doe", "John Smith"],
	"room": "Siebel 1440",
	"tags": ["BACKEND", "FRONTEND"],
	"number": 23
}
```

GET /project/
---------------------

Returns a list of all projects.

Request requires no body.

```json title="Example response"
{
	"projects": [
		{
			"id": "52fdfc072182654f163f5f0f9a621d72",
			"name": "Example Project 10",
			"description": "Example Project Description",
			"mentors": ["Jane Doe", "John Smith"],
			"room": "Siebel 1440",
			"tags": ["BACKEND", "FRONTEND"],
			"number": 23
		},
		{
			"id": "52fdfcab71282654f163f5f0f9a621d72",
			"name": "Example Project 11",
			"description": "Example Project Description",
			"mentors": ["Ann O. Nymous", "Joe Bloggs"],
			"room": "Siebel 1310",
			"tags": ["SYSTEMS"],
			"number": 33
		}
	]
}
```

GET /project/filter/?key=value
---------------------

Returns all projects, filtered with the given key-value pairs.

Request requires no body.

```json title="Example response"
{
	"projects": [
		{
			"id": "52fdfc072182654f163f5f0f9a621d72",
			"name": "Example Project 10",
			"description": "Example Project Description",
			"mentors": ["Jane Doe", "John Smith"],
			"room": "Siebel 1440",
			"tags": ["BACKEND", "FRONTEND"],
			"number": 23
		},
		{
			"id": "52fdfcab71282654f163f5f0f9a621d72",
			"name": "Example Project 11",
			"description": "Example Project Description",
			"mentors": ["Ann O. Nymous", "Joe Bloggs"],
			"room": "Siebel 1310",
			"tags": ["SYSTEMS"],
			"number": 33
		}
	]
}
```

POST /project/
-----------

Creates a project with the requested fields. Returns the created project.

```json title="Example request"
{
	"name": "Example Project 10",
	"description": "Example Project Description",
	"mentors": ["Jane Doe", "John Smith"],
	"room": "Siebel 1440",
	"tags": ["BACKEND", "FRONTEND"],
	"number": 23
}
```

```json title="Example response"
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"description": "Example Project Description",
	"mentors": ["Jane Doe", "John Smith"],
	"room": "Siebel 1440",
	"tags": ["BACKEND", "FRONTEND"],
	"number": 23
}
```

DELETE /project/PROJECTID/
-----------

Endpoint to delete a project with name `PROJECTID`.

Request requires no body.

```json title="Example response"
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"description": "Example Project Description",
	"mentors": ["Jane Doe", "John Smith"],
	"room": "Siebel 1440",
	"tags": ["BACKEND", "FRONTEND"],
	"number": 23
}
```

PUT /project/
----------

Updates the project with the id specified in the `id` field of the request. Returns the updated project.

```json title="Example request"
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"description": "Example Project Description",
	"mentors": ["Jane Doe", "John Smith"],
	"room": "Siebel 1440",
	"tags": ["BACKEND", "FRONTEND"],
	"number": 23
}
```

```json title="Example response"
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"description": "Example Project Description",
	"mentors": ["Jane Doe", "John Smith"],
	"room": "Siebel 1440",
	"tags": ["BACKEND", "FRONTEND"],
	"number": 23
}
```

GET /project/favorite/
--------------------

Returns the project favorites for the current user.

Request requires no body.

```json title="Example response"
{
	"id": "github001",
	"projects": [
		"52fdfc072182654f163f5f0f9a621d72",
		"34edfc072182654f163f5f0f9a621d72"
	]
}
```

POST /project/favorite/
-------------------------

Adds the given project to the favorites for the current user.

```json title="Example request"
{
	"projectId": "52fdfc072182654f163f5f0f9a621d72"
}
```

```json title="Example response"
{
	"id": "github001",
	"projects": [
		"52fdfc072182654f163f5f0f9a621d72",
		"34dffc072182654f163f5f0f9a621d72"
	]
}
```

DELETE /project/favorite/
----------------------------

Removes the given project from the favorites for the current user.

```json title="Example request"
{
	"projectId": "52fdfc072182654f163f5f0f9a621d72",
}
```

```json title="Example response"
{
	"id": "github001",
	"projects": [
		"52fdfc072182654f163f5f0f9a621d72"
	]
}
```
