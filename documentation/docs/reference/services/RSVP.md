RSVP
====

*Note:* The exact fields in the rsvp requests and responses will change based on the rsvp definitions provided in the API configuration file.

GET /rsvp/USERID/
-----------------

Returns the rsvp stored for the user with the `id` `USERID`.

Response format:
```
{
	"id": "github0000001"
	"isAttending": true,
}
```

GET /rsvp/
----------

Returns the rsvp stored for the user with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
	"id": "github0000001"
	"isAttending": true,
}
```

POST /rsvp/
-----------

Creates an rsvp for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
	"isAttending": true,
}
```

Response format:
```
{
	"id": "github0000001"
	"isAttending": true,
}
```

PUT /rsvp/
----------

Updated the rsvp for the user with the `id` in the JWT token provided in the Authorization header.

Request format:
```
{
	"isAttending": true,
}
```

Response format:
```
{
	"id": "github0000001"
	"isAttending": true,
}
```
