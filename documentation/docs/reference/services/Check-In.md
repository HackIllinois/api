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

Creates a checkin for the user with the associated user token `userToken` in the request body.

!!! note
    You will need a user token rather than a user id. User tokens are generated at the endpoints
    `GET /user/qr/` and `GET /user/qr/USERID/`. If no token is provided, then the request will
    `422`. If the token expires prior to when the request is received or the token is malformed,
    then a response with status `403` with the message `"Bad user token."`.

```json title="Example request"
{
	"userToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJnaXRodWIwMDAwMDAxIiwiZXhwIjoxNjc1Nzc1MjMzfQ.tREyQsEaG4TamXYZx8gNkY40-2FOdCr9n8dLrbk2UN8",
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
	"hasPickedUpSwag": true,
    "rsvpData": { ... }
}
```

PUT /checkin/
----------

Updates the checkin for the user with the associated user token `userToken` in the request body.

!!! note
    You will need a user token rather than a user id. User tokens are generated at the endpoints
    `GET /user/qr/` and `GET /user/qr/USERID/`. If no token is provided, then the request will
    `422`. If the token expires prior to when the request is received or the token is malformed,
    then a response with status `403` with the message `"Bad user token."`.

```json title="Example request"
{
	"userToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJnaXRodWIwMDAwMDAxIiwiZXhwIjoxNjc1Nzc1MjMzfQ.tREyQsEaG4TamXYZx8gNkY40-2FOdCr9n8dLrbk2UN8",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true
}
```

```json title="Example response"
{
	"id": "github0000001",
	"hasCheckedIn": true,
	"hasPickedUpSwag": true,
    "rsvpData": { ... }
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
