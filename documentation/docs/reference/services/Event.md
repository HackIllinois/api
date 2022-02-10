Event
=====

GET /event/EVENTID/
---------------------

Returns the event with the id of `EVENTID`.

Response format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Event 10",
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
			"id": "52fdfc072182654f163f5f0f9a621d72",
			"name": "Example Event 10",
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
			"name": "Example Event 11",
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

GET /event/filter/?key=value
---------------------

Returns all events, filtered with the given key-value pairs.

Response format:
```
{
    "events": [
        {
            "id": "52fdfc072182654f163f5f0f9a621d72",
            "name": "Example Event 10",
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
            "id": "9566c74d10037c4d7bbb0407d1e2c649",
            "name": "Example Event 10",
            "description": "This is a description",
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

POST /event/
-----------

Creates an event with the requested fields. Returns the created event.

Valid values for `eventType` are one of `MEAL SPEAKER WORKSHOP MINIEVENT QNA OTHER`.

Request format:
```
{
	"name": "Example Event 10",
	"description": "This is a description",
	"startTime": 1532202702,
	"endTime": 1532212702,
	"sponsor": "Example sponsor",
	"eventType": "WORKSHOP",
	"locations": [
		{
			"description": "Example Location",
			"tags": ["SIEBEL0", "ECEB1"],
			"latitude": 40.1138,
			"longitude": -88.2249
		}
	]
}
```

Response format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Event 10",
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
}
```

DELETE /event/EVENTID/
-----------

Endpoint to delete an event with name `EVENTID`. It removes the `EVENTID` from the event trackers, and every user's tracker.

Response format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Event 10",
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
}
```

PUT /event/
----------

Updates the event with the id specified in the `id` field of the request. Returns the updated event.

Request format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Event 10",
	"description": "This is a description",
	"startTime": 1532202702,
	"endTime": 1532212702,
	"sponsor": "Example sponsor",
	"eventType": "WORKSHOP",
	"locations": [
		{
			"description": "Example Location",
			"tags": ["SIEBEL0", "ECEB1"],
			"latitude": 40.1138,
			"longitude": -88.2249
		}
	]
}
```

Response format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"name": "Example Event 10",
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
}
```

POST /event/track/
------------------

Marks the specified user as attending the specified event. Returns the tracker for the user and the tracker for the event.

Request format:
```
{
	"eventId": "52fdfc072182654f163f5f0f9a621d72",
	"userId": "github0000001"
}
```

Response format:
```
{
	"eventTracker": {
		"eventId": "52fdfc072182654f163f5f0f9a621d72",
		"users": [
			"github0000001",
		]
	},
	"userTracker": {
		"userId": "github0000001",
		"events": [
			"52fdfc072182654f163f5f0f9a621d72"
		]
	}
}
```

GET /event/track/event/EVENTID/
---------------------------------

Returns the tracker for the event with the id `EVENTID`.

Response format:
```
{
	"eventId": "52fdfc072182654f163f5f0f9a621d72",
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
		"52fdfc072182654f163f5f0f9a621d72"
	]
}
```

GET /event/favorite/
--------------------

Returns the event favorites for the current user.

Response format:
```
{
	"id": "github001",
	"events": [
		"52fdfc072182654f163f5f0f9a621d72",
		"34edfc072182654f163f5f0f9a621d72"
	]
}
```

POST /event/favorite/
-------------------------

Adds the given event to the favorites for the current user.

Request format:
```
{
	"eventId": "52fdfc072182654f163f5f0f9a621d72"
}
```

Response format:
```
{
	"id": "github001",
	"events": [
		"52fdfc072182654f163f5f0f9a621d72",
		"34dffc072182654f163f5f0f9a621d72"
	]
}
```

DELETE /event/favorite/
----------------------------

Removes the given event from the favorites for the current user.

Request format:
```
{
	"eventId": "52fdfc072182654f163f5f0f9a621d72",
}
```

Response format:
```
{
	"id": "github001",
	"events": [
		"52fdfc072182654f163f5f0f9a621d72"
	]
}
```

GET /event/code/{id}/
----------------------------

Gets a struct that contains information about the event code (generated upon event creation) and expiration time.
By convention, event checkin codes will be 6 bytes long.

Response format:
```
{
    "id": "52fdfc072182654f163f5f0f9a621d72",
    "code": "sample_code",
    "expiration": 1521388800
}

```

POST /event/code/
----------------------------

Upserts a struct that contains information about the event code (generated upon event creation) and expiration time.

Request format:
```
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
    "code": "new_code",
    "expiration": 1521388800
}

```

Response format:
```
{
    "id": "52fdfc072182654f163f5f0f9a621d72",
    "code": "new_code",
    "expiration": 1521388800
}
```

POST /event/checkin/
----------------------------

Retrieves a struct that contains information about the event checkin status, point increment value, and total point number.
Takes in a struct that contains an event checkin code.

Valid values for `status` are `Success`, `InvalidCode`, `InvalidTime`, `AlreadyCheckedIn`. When `status != Success`, the `newPoints` and `totalPoints` fields will equal `-1` and should be ignored.

Request format:
```
{
    "code": "new_code",
}

```

Response format:
```
{
    "newPoints": 10,
    "totalPoints": 10,
    "status": "Success"
}
```