Decision
========

GET /decision/USERID/
----------------------------

Returns the decision stored for the user with the id `USERID`.

Request requires no body.

```json title="Example response"
{
	"finalized": false,
	"id": "github9279532",
	"status": "ACCEPTED",
	"wave": 1,
	"reviewer": "github9279532",
	"timestamp": 1526673862,
	"history": [
		{
			"finalized": false,
			"id": "github9279532",
			"status": "PENDING",
			"wave": 0,
			"reviewer": "github9279532",
			"timestamp": 1526673845
		},
		{
			"finalized": false,
			"id": "github9279532",
			"status": "ACCEPTED",
			"wave": 1,
			"reviewer": "github9279532",
			"timestamp": 1526673862
		}
	]
}
```

GET /decision/
----------------------------------

Returns the decision stored for the currently authenticated user (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
	"id": "github9279532",
	"status": "ACCEPTED"
}
```

POST /decision/
--------------------------

Updates the decision for the user as specified in the `id` field of the request. The full decision history is returned in the response.

```json title="Example request"
{
	"id": "github9279532",
	"status": "ACCEPTED",
	"wave": 1
}
```

```json title="Example response"
{
	"finalized": false,
	"id": "github9279532",
	"status": "ACCEPTED",
	"wave": 1,
	"reviewer": "github9279532",
	"timestamp": 1526673862,
	"history": [
		{
			"finalized": false,
			"id": "github9279532",
			"status": "PENDING",
			"wave": 0,
			"reviewer": "github9279532",
			"timestamp": 1526673845
		},
		{
			"finalized": false,
			"id": "github9279532",
			"status": "ACCEPTED",
			"wave": 1,
			"reviewer": "github9279532",
			"timestamp": 1526673862
		},
		{
			"finalized": true,
			"id": "github9279532",
			"status": "ACCEPTED",
			"wave": 1,
			"reviewer": "github9279532",
			"timestamp": 1526673862
		}
	]
}
```

POST /decision/finalize/
--------------------------

Finalizes / unfinalizes the decision for the current user. The full decision history is returned in the response. This endpoint will return an AttributeMismatchError if the requested action results in a Finalized status matching the current Finalized status. 

```json title="Example request"
{
	"id": "github9279532",
	"finalized": true
}
```

```json title="Example response"
{
	"finalized": true,
	"id": "github9279532",
	"status": "ACCEPTED",
	"wave": 1,
	"reviewer": "github9279532",
	"timestamp": 1526673862,
	"history": [
		{
			"finalized": false,
			"id": "github9279532",
			"status": "PENDING",
			"wave": 0,
			"reviewer": "github9279532",
			"timestamp": 1526673845
		},
		{
			"finalized": true,
			"id": "github9279532",
			"status": "ACCEPTED",
			"wave": 1,
			"reviewer": "github9279532",
			"timestamp": 1526673862
		}
	]

}
```

GET /decision/filter/?key=value
----------------------------------

Returns the user decisions, filtered with the given key-value pairs.

Request requires no body.

```json title="Example response"
{
	"decisions": [
		{
			"finalized": false,
			"id": "github9279532",
			"status": "ACCEPTED",
			"wave": 1,
			"reviewer": "github9279532",
			"timestamp": 1526673862,
			"history": [
				{
					"finalized": false,
					"id": "github9279532",
					"status": "PENDING",
					"wave": 0,
					"reviewer": "github9279532",
					"timestamp": 1526673845
				},
				{
					"finalized": false,
					"id": "github9279532",
					"status": "ACCEPTED",
					"wave": 1,
					"reviewer": "github9279532",
					"timestamp": 1526673862
				}
			]
		},
		{
			"finalized": false,
			"id": "github9279533",
			"status": "ACCEPTED",
			"wave": 1,
			"reviewer": "github9279533",
			"timestamp": 1526673863,
			"history": [
				{
					"finalized": false,
					"id": "github9279533",
					"status": "PENDING",
					"wave": 0,
					"reviewer": "github9279533",
					"timestamp": 1526673846
				},
				{
					"finalized": false,
					"id": "github9279533",
					"status": "ACCEPTED",
					"wave": 1,
					"reviewer": "github9279533",
					"timestamp": 1526673863
				}
			]
		}
	]
}
```
