Authorization
=============

GET /auth/?provider=PROVIDER
----------------------------

Redirects to the `PROVIDER`'s oauth authorization page. Once the user accepts the oauth authorization they will be redirected to the client's auth page with an oauth code. This code should be sent to the api to be exchanged for an api jwt.

Valid `PROVIDER` strings: `github`

POST /auth/code/?provider=PROVIDER
----------------------------------

Exchanges a valid oauth code from a jwt from the api. This jwt should be placed in the `Authorization` header for all future api requests.

Valid `PROVIDER` strings: `github`

Request format:
```
{
	"code": "5897dk3j05192c5j2gc8"
}
```

Response format:
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFybmF2c2Fua2FyYW5AZ21haWwuY29tIiwiZXhwIjoxNTI1ODQ1MzA0LCJpZCI6MCwicm9sZXMiOlsiVXNlciJdfQ.lYxFGSNDU9q7FoQHNHGvpKu1fTHf8yHsKPg8FDt9L-s"
}
```
GET /auth/roles/?id=USERID
--------------------------

Gets the roles of the user with the id `USERID`.

Response format:
```
{
	"id": "github6892396",
	"roles": [
		"User"
	]
}
```

PUT /auth/roles/
-----------------

Sets the roles of the user with the given `id` to `roles`. The updated user's roles will be returned.

Request format:
```
{
	"id": "github6892396",
	"roles": [
		"User"
	]
}
```

Response format:
```
{
	"id": "github6892396",
	"roles": [
		"User"
	]
}
```
