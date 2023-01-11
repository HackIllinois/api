Checkin
=======

GET /checkin/USERID/
-----------------

Returns the checkin stored for the user with the `id` `USERID`.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

GET /checkin/
----------

Returns the checkin stored for the currently authenticated user (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

POST /checkin/
-----------

Creates an checkin for the user with the `id` in the request body.

```json title="Example request"
{
	"id": "github0000001",
	"override": true,
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

```json title="Example response"
{
	"id": "github0000001",
	"override": true,
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

PUT /checkin/
----------

Updated the checkin for the user with the `id` in the request body.

```json title="Example request"
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

```json title="Example response"
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

GET /checkin/list/
----------

Returns a list of all user IDs for users who have checked in

Request requires no body.

```json title="Example response"
{
	"checkedInUsers": [
		"github0000001",
		"github0000005"
	]
}
```
