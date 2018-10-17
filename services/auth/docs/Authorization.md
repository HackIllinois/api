Authorization
=============

GET /auth/PROVIDER/?redirect_uri=AUTHREDIRECTURI
------------------------------------------------

Redirects to the `PROVIDER`'s OAuth authorization page. Once the user accepts the OAuth authorization they will be redirected to the client's auth page with an OAuth code. This code should be sent to the API to be exchanged for an API JWT.

Valid `PROVIDER` strings: `github`, `google`

`AUTHREDIRECTURI` can be specified to override the default OAuth redirect URI. This is the URI to which the application is redirected after the Authorization request is approved / rejected.

POST /auth/code/PROVIDER/?redirect_uri=AUTHREDIRECTURI
------------------------------------------------------

Exchanges a valid OAuth code from a JWT from the API. This JWT should be placed in the `Authorization` header for all future API requests.

Valid `PROVIDER` strings: `github`, `google` and `linkedin`.

`AUTHREDIRECTURI` can be specified to override the default OAuth redirect URI. This is the URI to which the application is redirected after the token request is completed.

*Important note:* For Google OAuth requests, the provided `redirect_uri` needs to be the same as the one provided in the GET request above. If the two `redirect_uri`s differ, Google will reject the OAuth token request.

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
GET /auth/roles/USERID/
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

GET /auth/token/refresh/
-----------------

Creates a new JWT for the current user. This is useful when the user's roles change, and the updated roles need to be encoded into a new JWT, such as during registration. 

Response format:
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFybmF2c2Fua2FyYW5AZ21haWwuY29tIiwiZXhwIjoxNTI1ODQ1MzA0LCJpZCI6MCwicm9sZXMiOlsiVXNlciJdfQ.lYxFGSNDU9q7FoQHNHGvpKu1fTHf8yHsKPg8FDt9L-s"
}
```
