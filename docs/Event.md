Event
=====

GET /event/EVENTNAME/
---------------------

Returns the event with the name of `EVENTNAME`. `EVENTNAME` should be url encoded.

Response format:
```
{
	"name": "Example Event",
	"description": "This is a description",
	"startTime": 1532202702,
	"endTime": 1532212702,
	"locationDescription": "Example Location",
	"latitude": 40.1138,
	"longitude": -88.2249,
	"sponsor": "Example sponsor",
	"eventType": "WORKSHOP"
}
```

GET /event/
---------------------

Returns a list of all events.

Response format:
```
{
	events: [
		{
			"name": "Example Event",
			"description": "This is a description",
			"startTime": 1532202702,
			"endTime": 1532212702,
			"locationDescription": "Example Location",
			"latitude": 40.1138,
			"longitude": -88.2249,
			"sponsor": "Example sponsor",
			"eventType": "WORKSHOP"
		},
		{
			"name": "Example Event 2",
			"description": "This is another description",
			"startTime": 1532202703,
			"endTime": 1532212703,
			"locationDescription": "Example Location 2",
			"latitude": 40.1139,
			"longitude": -88.2250,
			"sponsor": "Example sponsor 2",
			"eventType": "WORKSHOP"
		}
	]
}
```

POST /event/
-----------

Creates an event with the requested fields. Returns the created event.

Request format:
```
{
	"name": "Example Event",
	"description": "This is a description",
	"startTime": 1532202702,
	"endTime": 1532212702,
	"locationDescription": "Example Location",
	"latitude": 40.1138,
	"longitude": -88.2249,
	"sponsor": "Example sponsor",
	"eventType": "WORKSHOP"
}
```

Response format:
```
{
	"name": "Example Event",
	"description": "This is a description",
	"startTime": 1532202702,
	"endTime": 1532212702,
	"locationDescription": "Example Location",
	"latitude": 40.1138,
	"longitude": -88.2249,
	"sponsor": "Example sponsor",
	"eventType": "WORKSHOP"
}
```

PUT /event/
----------

Updates the event with the name specified in the `name` field of the request. Returns the updated event.

Request format:
```
{
	"name": "Example Event",
	"description": "This is an updated description",
	"startTime": 1532202702,
	"endTime": 1532212702,
	"locationDescription": "Example Location",
	"latitude": 40.1138,
	"longitude": -88.2249,
	"sponsor": "Example sponsor",
	"eventType": "WORKSHOP"
}
```

Response format:
```
{
	"name": "Example Event",
	"description": "This is an updated description",
	"startTime": 1532202702,
	"endTime": 1532212702,
	"locationDescription": "Example Location",
	"latitude": 40.1138,
	"longitude": -88.2249,
	"sponsor": "Example sponsor",
	"eventType": "WORKSHOP"
}
```

POST /event/track/
------------------

Marks the specified user as attending the specified event. Returns the tracker for the user and the tracker for the event.

Request format:
```
{
	"eventName": "Example Event",
	"userId": "github0000001"
}
```

Response format:
```
{
	"eventTracker": {
		"eventName": "Example Event",
		"users": [
			"github0000001",
		]
	},
	"userTracker": {
		"userId": "github0000001",
		"events": [
			"Example Event"
		]
	}
}
```

GET /event/track/event/EVENTNAME/
---------------------------------

Returns the tracker for the event with the name `EVENTNAME`.

Response format:
```
{
	"eventName": "Example Event",
	"users": [
		"github0000001",
	]
}
```

GET /event/track/user/USERID/
-----------------------------

Returns the tracker for the user with the id `USERID`.

Response format:
```
{
	"userId": "github0000001",
	"events": [
		"Example Event"
	]
}
```
