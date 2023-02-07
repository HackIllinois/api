Event
=====

The `isAsync` field for events is optional. If it is not specified or is false, then `startTime` and
`endTime` are required. Otherwise, `startTime` and `endTime` are optional.

The fields `isPrivate` and `displayOnStaffCheckin` are private are only visible to users that have
the Staff or Admin role.

```json title="Example struct a non-staff/non-admin will receive"
{
    "id": "93d91d48a5b111edafa10242ac120002",
	"name": "Example Event 1",
	"description": "This is a placeholder description",
	"startTime": 1532202702,
	"endTime": 1532212702,
    "sponsor": "",
	"eventType": "MEAL",
    "locations": [
		{
	        "description": "Location info here",
			"tags": ["SIEBEL3", "CIF"],
			"latitude":    123.456,
			"longitude":   123.456,
		},
	"points": 0,
}
```

```json title="Example struct a staff/admin will receive"
{
    "id": "93d91d48a5b111edafa10242ac120002",
	"name": "Example Event 1",
	"description": "This is a placeholder description",
	"startTime": 1532202702,
	"endTime": 1532212702,
    "sponsor": "",
	"eventType": "MEAL",
    "locations": [
		{
	        "description": "Location info here",
			"tags": ["SIEBEL3", "CIF"],
			"latitude":    123.456,
			"longitude":   123.456,
		},
	"points": 0,
	"isPrivate": false,
	"displayOnStaffCheckin": true,
}
```

GET /event/EVENTID/
---------------------

Returns the event with the id of `EVENTID`.

Request requires no body.

```json title="Example response"
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

Request requires no body.

```json title="Example response"
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

Request requires no body.

```json title="Example response"
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

Valid values for `eventType` are `MEAL`, `SPEAKER`, `WORKSHOP`, `MINIEVENT`, `QNA`, or `OTHER`.

```json title="Example request"
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
	"isAsync": false
}
```

```json title="Example response"
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

Request requires no body.

```json title="Example response"
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

```json title="Example request"
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

```json title="Example response"
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

```json title="Example request"
{
	"eventId": "52fdfc072182654f163f5f0f9a621d72",
	"userId": "github0000001"
}
```

```json title="Example response"
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

Request requires no body.

```json title="Example response"
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

Request requires no body.

```json title="Example response"
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

Request requires no body.

```json title="Example response"
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

```json title="Example request"
{
	"eventId": "52fdfc072182654f163f5f0f9a621d72"
}
```

```json title="Example response"
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

```json title="Example request"
{
	"eventId": "52fdfc072182654f163f5f0f9a621d72",
}
```

```json title="Example response"
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

Request requires no body.

```json title="Example response"
{
    "id": "52fdfc072182654f163f5f0f9a621d72",
    "code": "sample_code",
    "expiration": 1521388800
}

```

PUT /event/code/{id}/
----------------------------

Updates a struct that contains information about the event code (generated upon event creation) and expiration time.

```json title="Example request"
{
    "code": "new_code",
    "expiration": 1521388800
}

```

```json title="Example response"
{
    "id": "52fdfc072182654f163f5f0f9a621d72",
    "code": "new_code",
    "expiration": 1521388800
}
```

POST /event/staff/checkin/
----------------------------

Used for staff to check in attendees to various events.

Request should include an attendee's user token (`userToken`) and an event id (`eventId`).

Returns a status, the user's new points, and the user's total points.

Valid values for `status` are `Success`, `InvalidEventId`, `BadUserToken`, `AlreadyCheckedIn`. 

!!! warning
	Please be aware that `BadUserToken` will be returned if the user token is **invalid, expired, or malformed**. 

	If you get this status, the staff should see a message like `Bad user token. Try asking the attendee to refresh their qr code.`
	To do this, the attendee will need to request `GET /user/qr/` again.

!!! note
	When `status != Success`, the `newPoints` and `totalPoints` fields will equal `-1` and should be ignored.

!!! note
    On status `Sucess` and `AlreadyCheckedIn`, the field `rsvpData` will be populated with the
    user's RSVP data and registration data.

!!! note
	The user token `some_token` should be retrieved from the `userToken` field of a user QR code URI
    (`hackillinois://user?userToken=some_token`)

```json title="Example request"
{
	"userToken": "some_token",
	"eventId": "some_event_id"
}
```

```json title="Example response"
{
    "newPoints": 10,
    "totalPoints": 10,
    "status": "Success",
    "rsvpData": {
        "id": "github0123456",
        "isAttending": true,
        "registrationData": { ... }
    }
}
```

POST /event/checkin/
----------------------------

Used for attendees to check into various events. Like `/event/staff/checkin/`, but doesn't require staff verification.

Request should include an event code (`code`).

Returns a status, the user's new points, and the user's total points.

Valid values for `status` are `Success`, `InvalidCode`, `ExpiredOrProspective`, `AlreadyCheckedIn`. 

!!! note
	`ExpiredOrProspective` in this case means the event has already happened, or has not started yet.

!!! note
	When `status != Success`, the `newPoints` and `totalPoints` fields will equal `-1` and should be ignored.

!!! note
	The `code` should be retrieved from the `code` field of a event QR code URI (`hackillinois://event?code=some_event_code`)

```json title="Example request"
{
    "code": "some_event_code"
}
```

```json title="Example response"
{
    "newPoints": 10,
    "totalPoints": 10,
    "status": "Success"
}
```
