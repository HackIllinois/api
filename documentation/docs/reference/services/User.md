Users
=====

GET /user/USERID/
-----------------

Returns the basic user information stored for the user with the `id` `USERID`.

Response format:
```
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

Returns the basic user information stored for the user associated with the `id` in the given JWT in the Authorization header.

Response format:
```
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

Request format:
```
{
	"id": "github000001",
	"username": "test",
	"firstName": "ExampleFirstName",
	"lastName": "ExampleLastName",
	"email": "test@gmail.com"
}
```

Response format:
```
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

Response format:
```
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
