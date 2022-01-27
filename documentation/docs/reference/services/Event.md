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
	"eventType": "WORKSHOP",
	"inPersonPoints": 10,
	"inPersonVirtPoints": 5,
	"virtualPoints": 3
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
			"eventType": "WORKSHOP",
			"inPersonPoints": 10,
			"inPersonVirtPoints": 5,
			"virtualPoints": 3
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
			"eventType": "WORKSHOP",
			"inPersonPoints": 20,
			"inPersonVirtPoints": 10,
			"virtualPoints": 5
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
            "eventType": "WORKSHOP",
			"inPersonPoints": 10,
			"inPersonVirtPoints": 5,
			"virtualPoints": 3
        },
        {
            "id": "9566c74d10037c4d7bbb0407d1e2c649",
            "name": "Example Event 11",
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
            "eventType": "WORKSHOP",
			"inPersonPoints": 20,
			"inPersonVirtPoints": 10,
			"virtualPoints": 5
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
	],
	"inPersonPoints": 10,
	"inPersonVirtPoints": 5,
	"virtualPoints": 3
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
	"eventType": "WORKSHOP",
	"inPersonPoints": 10,
	"inPersonVirtPoints": 5,
	"virtualPoints": 3
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
	"eventType": "WORKSHOP",
	"inPersonPoints": 10,
	"inPersonVirtPoints": 7,
	"virtualPoints": 3
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
	],
	"inPersonPoints": 20,
	"inPersonVirtPoints": 10,
	"virtualPoints": 5
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
	"eventType": "WORKSHOP",
	"inPersonPoints": 20,
	"inPersonVirtPoints": 10,
	"virtualPoints": 5
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

Gets an array of structs that contains information about the event codes (generated upon event creation) and expiration times.
By convention, event checkin codes will be 6 bytes long.

Response format:
```
[
	{
		"codeID": "sample_code_1",
		"eventID": "52fdfc072182654f163f5f0f9a621d72",
		"isVirtual": false,
		"expiration": 1521388800
	},
	{
		"codeID": "sample_code_2",
		"eventID": "52fdfc072182654f163f5f0f9a621d72",
		"isVirtual": true,
		"expiration": 1521388800
	}
]
```

POST /event/code/
----------------------------

Upserts a struct that contains information about the event code (generated upon event creation) and expiration time.

NOTE: Once created, the code ID cannot be changed.

Request format:
```
{
	"codeID": "code",
	"eventID": "52fdfc072182654f163f5f0f9a621d72",
	"isVirtual": true,
	"expiration": 1521388800
}

```

Response format:
```
{
	"codeID": "code",
	"eventID": "52fdfc072182654f163f5f0f9a621d72",
	"isVirtual": true,
	"expiration": 1521388800
}
```

POST /event/checkin/
----------------------------

Retrieves a struct that contains information about the event checkin status, point increment value, and total point number.
Takes in a struct that contains an event checkin code.

The endpoint will check the use HackIllinois-Identity field to determine if the user redeeming the points registered for virtual or in-person. It will compare this to the scanned code and will award points accordingly. (e.g. If registered for virtual, give `virtualPoints`. If registered for in-person but attended virtually, give `inPersonVirtPoints`. If registered for in-person and attended in-person, give `inPersonPoints`.)

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