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
