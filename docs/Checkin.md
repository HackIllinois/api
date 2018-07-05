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

GET /checkin/qr/
----------

Get the string to be embedded in the current user's QR code. 
The QR code string will contain information stored in the form of a URI.

Response format:
```
{
	"id": "github0000001",
	"qrInfo": "hackillinois://info?name=FirstName&hasCheckedIn=true&hasPickedUpSwag=true"
}
```

GET /checkin/qr/{id}/
----------

Get the string to be embedded in the specified user's QR code. 
The QR code string will contain information stored in the form of a URI.

Response format:
```
{
	"id": "github0000001",
	"qrInfo": "hackillinois://info?name=FirstName&hasCheckedIn=true&hasPickedUpSwag=true"
}
```
