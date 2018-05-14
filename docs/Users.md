Users
=====

GET /user/USERID/
----------------------------

Returns the basic user information stored for the user with the `id` `USERID`.

Response format:
```
{
	"id": "github09829234",
	"username": "ExampleUsername",
	"email": "examplemail@gmail.com"
}
```

GET /user/
----------------------------------

Returns the basic user information stored for the user associated with the `id` in the given JWT in the Authorization header.

Response format:
```
{
	"id": "github09829234",
	"username": "ExampleUsername",
	"email": "examplemail@gmail.com"
}
```

POST /user/
--------------------------

Sets the basic user information for the user as specified in the `id` field of the request. The information recorded in the database is returned in the response.

Request format:
```
{
	"id": "github000001",
	"username": "test",
	"email": "test@gmail.com"
}
```

Response format:
```
{
	"id": "github000001",
	"username": "test",	
	"email": "test@gmail.com"
}
```
