Users
=====

GET /user/USERID/
-----------------

Returns the basic user information stored for the user with the `id` `USERID`.

Request requires no body.

```json title="Example response"
{
	"id": "github09829234",
	"username": "ExampleUsername",
	"firstName": "ExampleFirstName",
	"lastName": "ExampleLastName",
	"email": "examplemail@gmail.com"
}
```

GET /user/
----------

Returns the basic user information stored for the currently authenticated user (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
	"id": "github09829234",
	"username": "ExampleUsername",
	"firstName": "ExampleFirstName",
	"lastName": "ExampleLastName",
	"email": "examplemail@gmail.com"
}
```

POST /user/
-----------

Sets the basic user information for the user as specified in the `id` field of the request. The information recorded in the database is returned in the response.

```json title="Example request"
{
	"id": "github000001",
	"username": "test",
	"firstName": "ExampleFirstName",
	"lastName": "ExampleLastName",
	"email": "test@gmail.com"
}
```

```json title="Example response"
{
	"id": "github000001",
	"username": "test",
	"firstName": "ExampleFirstName",
	"lastName": "ExampleLastName",
	"email": "test@gmail.com"
}
```

GET /user/filter/?key=value
---------------------------

Returns the basic user information, filtered with the given key-value pairs.

To paginate the response, provide a parameter "p" with the page number you are requesting, as well as a parameter "limit" with the desired number of Users per page.
If the pagination request exceeds the length of the available Users, it will be truncated. 
For example, the following request: `/user/filter/?key=value&p=1&limit=5` will return the first 5 Users (index 0 through 4).

To sort the users, provide a **comma-separated** "sortby" parameter. 
For example, the following request: `/user/filter/?key=value&sortby=FirstName,LastName` will return a list of filtered users sorted by first name, using the last name as a tie breaker.

To reverse the sort, add a minus (-) to the desired sort field. For example, `FirstName` would become `-FirstName`.

Request requires no body.

```json title="Example response"
{
	"users": [
		{
			"id": "github09829234",
			"username": "ExampleUsername",
			"firstName": "ExampleFirstName",
			"lastName": "ExampleLastName",
			"email": "examplemail@gmail.com"
		},
		{
			"id": "github09829235",
			"username": "ExampleUsername2",
			"firstName": "ExampleFirstName2",
			"lastName": "ExampleLastName2",
			"email": "examplemail2@gmail.com"
		}
	]
}
```

GET /user/qr/
----------

Get the string to be embedded in the current user's QR code. 
The QR code string will contain information stored in the form of a URI.

!!! warning
	The URI wll contain a JWT that will expire in 20 seconds after the request was received. If you
	need a new user token, poll this endpoint again and it will give a fresh token.
	
	Since each token expires after 20 seconds, it is recommended that you poll every 15 seconds to 
	allow for any last second user QR scans to succeed.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"qrInfo": "hackillinois://user?userToken=mWZfc9b4zLEzyqqqFmSbvFcEXuY2CEjiAqWhbrVdzcc"
}
```

GET /user/qr/{id}/
----------

Get the string to be embedded in the specified user's QR code. 
The QR code string will contain information stored in the form of a URI.

!!! warning
	See `GET /user/qr/` for more information the user token's lifetime that is embedded into the URI.
	If you need a new user token, poll this endpoint again and it will give a fresh token.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"qrInfo": "hackillinois://user?userToken=mWZfc9b4zLEzyqqqFmSbvFcEXuY2CEjiAqWhbrVdzcc"
}
```
