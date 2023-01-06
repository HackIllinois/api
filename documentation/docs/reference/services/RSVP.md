RSVP
====

!!! warning
	The exact fields in the rsvp requests and responses **will change** based on the rsvp definitions provided in the API configuration file.
	Please consult them accordingly.

GET /rsvp/USERID/
-----------------

Returns the rsvp stored for the user with the `id` `USERID`.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"isAttending": true
}
```

GET /rsvp/
----------

Returns the rsvp stored for the currently authenticated user (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"isAttending": true
}
```

POST /rsvp/
-----------

Creates an rsvp for the currently authenticated user (determined by the JWT in the `Authorization` header).

```json title="Example request"
{
	"isAttending": true
}
```

```json title="Example response"
{
	"id": "github0000001",
	"isAttending": true
}
```

PUT /rsvp/
----------

Updated the rsvp for the currently authenticated user (determined by the JWT in the `Authorization` header).

```json title="Example request"
{
	"isAttending": true
}
```

```json title="Example response"
{
	"id": "github0000001",
	"isAttending": true
}
```
