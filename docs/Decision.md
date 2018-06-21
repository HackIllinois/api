Decision
========

GET /decision/USERID/
----------------------------

Returns the decision stored for the user with the `id` `USERID`.

Response format:
```
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

Returns the decision stored for the user associated with the `id` in the given JWT in the Authorization header.

Response format:
```
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

POST /decision/
--------------------------

Updates the decision for the user as specified in the `id` field of the request. The full decision history is returned in the response. 

Request format:
```
{
	"finalized": false,
	"id": "github9279532",
	"status": "ACCEPTED",
	"wave": 1
}
```

Response format:
```
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

Finalizes the decision for the user as specified in the `id` field of the request. The full decision history is returned in the response. 

Request format:
```
{
	"finalized": false,
	"id": "github9279532",
	"status": "ACCEPTED",
	"wave": 1
}
```

Response format:
```
{
	"finalized": true,
	"id": "github9279532",
	"status": "ACCEPTED",
	"wave": 1,
	"reviewer": "github9279532",
	"timestamp": 1526673862,
	"history": [
		{
			"id": "github9279532",
			"status": "PENDING",
			"wave": 0,
			"reviewer": "github9279532",
			"timestamp": 1526673845
		},
		{
			"id": "github9279532",
			"status": "ACCEPTED",
			"wave": 1,
			"reviewer": "github9279532",
			"timestamp": 1526673862
		}
	]
}
```