Checkin
=======

GET /checkin/USERID/
-----------------

Returns the checkin stored for the user with the `id` `USERID`.

Response format:
```
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

GET /checkin/
----------

Returns the checkin stored for the user with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

POST /checkin/
-----------

Creates an checkin for the user with the `id` in the request body.

Request format:
```
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

Response format:
```
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

PUT /checkin/
----------

Updated the checkin for the user with the `id` in the request body.

Request format:
```
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

Response format:
```
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```
