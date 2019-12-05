Project
=====

GET /project/PROJECTID/
---------------------

Returns the project with the id of `PROJECTID`.

Response format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"mentors": ["Jane Doe", "John Smith"],
	"location": {
		"description": "Example Location",
		"latitude": 40.1138,
		"longitude": -88.2249
	}
	"tags": ["BACKEND", "FRONTEND"],
	"code": "A1"
}
```

GET /project/
---------------------

Returns a list of all projects.

Response format:
```
{
	"projects": [
		{
			"id": "52fdfc072182654f163f5f0f9a621d72",
			"name": "Example Project 10",
			"mentors": ["Jane Doe", "John Smith"],
			"location": {
				"description": "Example Location",
				"latitude": 40.1138,
				"longitude": -88.2249
			}
			"tags": ["BACKEND", "FRONTEND"],
			"code": "A1"
		},
		{
			"id": "52fdfcab71282654f163f5f0f9a621d72",
			"name": "Example Project 11",
			"mentors": ["Ann O. Nymous", "Joe Bloggs"],
			"location": {
				"description": "Example Location",
				"latitude": 77.1238,
				"longitude": -84.3249
			}
			"tags": ["SYSTEMS"],
			"code": "B2"
		}
	]
}
```

GET /project/filter/?key=value
---------------------

Returns all projects, filtered with the given key-value pairs.

Response format:
```
{
	"projects": [
		{
			"id": "52fdfc072182654f163f5f0f9a621d72",
			"name": "Example Project 10",
			"mentors": ["Jane Doe", "John Smith"],
			"location": {
				"description": "Example Location",
				"latitude": 40.1138,
				"longitude": -88.2249
			}
			"tags": ["BACKEND", "FRONTEND"],
			"code": "A1"
		},
		{
			"id": "52fdfcab71282654f163f5f0f9a621d72",
			"name": "Example Project 11",
			"mentors": ["Ann O. Nymous", "Joe Bloggs"],
			"location": {
				"description": "Example Location",
				"latitude": 77.1238,
				"longitude": -84.3249
			}
			"tags": ["SYSTEMS"],
			"code": "B2"
		}
	]
}
```

POST /project/
-----------

Creates a project with the requested fields. Returns the created project.

Request format:
```
{
	"name": "Example Project 10",
	"mentors": ["Jane Doe", "John Smith"],
	"location": {
		"description": "Example Location",
		"latitude": 40.1138,
		"longitude": -88.2249
	}
	"tags": ["BACKEND", "FRONTEND"],
	"code": "A1"
}
```

Response format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"mentors": ["Jane Doe", "John Smith"],
	"location": {
		"description": "Example Location",
		"latitude": 40.1138,
		"longitude": -88.2249
	}
	"tags": ["BACKEND", "FRONTEND"],
	"code": "A1"
}
```

DELETE /project/PROJECTID/
-----------

Endpoint to delete a project with name `PROJECTID`.

Response format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"mentors": ["Jane Doe", "John Smith"],
	"location": {
		"description": "Example Location",
		"latitude": 40.1138,
		"longitude": -88.2249
	}
	"tags": ["BACKEND", "FRONTEND"],
	"code": "A1"
}
```

PUT /project/
----------

Updates the project with the id specified in the `id` field of the request. Returns the updated project.

Request format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"mentors": ["Jane Doe", "John Smith"],
	"location": {
		"description": "Example Location",
		"latitude": 40.1138,
		"longitude": -88.2249
	}
	"tags": ["BACKEND", "FRONTEND"],
	"code": "A1"
}
```

Response format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Project 10",
	"mentors": ["Jane Doe", "John Smith"],
	"location": {
		"description": "Example Location",
		"latitude": 40.1138,
		"longitude": -88.2249
	}
	"tags": ["BACKEND", "FRONTEND"],
	"code": "A1"
}
```
