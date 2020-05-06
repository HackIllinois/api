Recognition
=====


GET /recognition/
---------------------

Returns a list of all events.

Response format:
```
{
	events: [
		{
			"id": "52fdfc072182654f163f5f0f9a621d72",
			"name": "Example Recognition 10",
			"description": "This is a description",
			"startTime": 1532202702,
			"endTime": 1532212702,
			"locations": [
				{
					"description": "Example Location",
					"tags": ["SIEBEL0", "ECEB1"],
					"latitude": 40.1138,
					"longitude": -88.2249
				}
			],
			"sponsor": "Example sponsor",
			"eventType": "WORKSHOP"
		},
		{
			"id": "52fdfcab71282654f163f5f0f9a621d72",
			"name": "Example Recognition 11",
			"description": "This is another description",
			"startTime": 1532202702,
			"endTime": 1532212702,
			"locations": [
				{
					"description": "Example Location",
					"tags": ["SIEBEL3"],
					"latitude": 40.1138,
					"longitude": -88.2249
				}
			],
			"sponsor": "Example sponsor",
			"eventType": "WORKSHOP"
		}
	]
}
```

POST /recognition/
-----------

Creates an event with the requested fields. Returns the created event.

Request format:
```
{
	"name": "Example Recognition 10",
	"description": "This is a description",
	"presenter": "Example presenter",
	"eventId": "81855ad8681d0d86d1e91e00167939cb",
	"tags": ["Data Science", "Mobile"],
	"recipients": [
		{
			"type": "ALL"
		}
	]
}
```

Response format:
```
{
	"name": "Example Recognition 10",
	"description": "This is a description",
	"presenter": "Example presenter",
	"eventId": "81855ad8681d0d86d1e91e00167939cb",
	"tags": ["Data Science", "Mobile"],
	"recipients": [
		{
			"type": "ALL"
		}
	]
}
```

DELETE /recognition/RECOGNITIONID/
-----------

Endpoint to delete an event with name `RECOGNITIONID`

Response format:
```
{
	"name": "Example Recognition 10",
	"description": "This is a description",
	"presenter": "Example presenter",
	"eventId": "81855ad8681d0d86d1e91e00167939cb",
	"tags": ["Data Science", "Mobile"],
	"recipients": [
		{
			"type": "ALL"
		}
	]
}
```
