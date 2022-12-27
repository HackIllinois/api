Authorization
=============

GET /auth/PROVIDER/?redirect_uri=AUTHREDIRECTURI
------------------------------------------------

Redirects to the `PROVIDER`'s OAuth authorization page. Once the user accepts the OAuth authorization they will be redirected to the client's auth page with an OAuth code. This code should be sent to the API to be exchanged for an API JWT.

Valid `PROVIDER` strings: `github`, `google`, `linkedin`

`AUTHREDIRECTURI` can be specified to override the default OAuth redirect URI. This is the URI to which the application is redirected after the Authorization request is approved / rejected.

POST /auth/code/PROVIDER/?redirect_uri=AUTHREDIRECTURI
------------------------------------------------------

Exchanges a valid OAuth code from a JWT from the API. This JWT should be placed in the `Authorization` header for all future API requests.

Valid `PROVIDER` strings: `github`, `google` and `linkedin`.

`AUTHREDIRECTURI` can be specified to override the default OAuth redirect URI. This is the URI to which the application is redirected after the token request is completed.

!!! warning
	For Google OAuth requests, the provided `redirect_uri` needs to be the same as the one provided in the GET request above.
	If the two `redirect_uri`s differ, Google will reject the OAuth token request.

```json title="Example request"
{
	"code": "5897dk3j05192c5j2gc8"
}
```

```json title="Example response"
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFybmF2c2Fua2FyYW5AZ21haWwuY29tIiwiZXhwIjoxNTI1ODQ1MzA0LCJpZCI6MCwicm9sZXMiOlsiVXNlciJdfQ.lYxFGSNDU9q7FoQHNHGvpKu1fTHf8yHsKPg8FDt9L-s"
}
```

GET /auth/roles/USERID/
-----------------------

Gets the roles of the user with the id `USERID`.

Request requires no body.

```json title="Example response"
{
	"id": "github6892396",
	"roles": [
		"User"
	]
}
```

GET /auth/roles/
-----------------------

Gets the roles of the currently authenticated user (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
	"id": "github6892396",
	"roles": [
		"User"
	]
}
```

PUT /auth/roles/add/
--------------------

Adds the given `role` to the user with the given `id`. The updated user's roles will be returned.

```json title="Example request"
{
	"id": "github6892396",
	"role": "User"
}
```

```json title="Example response"
{
	"id": "github6892396",
	"roles": [
		"User"
	]
}
```

PUT /auth/roles/remove/
-----------------------

Removes the given `role` from the user with the given `id`. The updated user's roles will be returned.

```json title="Example request"
{
	"id": "github6892396",
	"role": "User"
}
```

```json title="Example response"
{
	"id": "github6892396",
	"roles": []
}
```

GET /auth/token/refresh/
------------------------

Creates a new JWT for the current user. This is useful when the user's roles change, and the updated roles need to be encoded into a new JWT, such as during registration.

Request requires no body.

```json title="Example response"
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFybmF2c2Fua2FyYW5AZ21haWwuY29tIiwiZXhwIjoxNTI1ODQ1MzA0LCJpZCI6MCwicm9sZXMiOlsiVXNlciJdfQ.lYxFGSNDU9q7FoQHNHGvpKu1fTHf8yHsKPg8FDt9L-s"
}
```

GET /auth/roles/list/
-----------------------

Gets the list of valid roles a user can have.

Request requires no body.

```json title="Example response"
{
	"roles": [
		"Admin",
		"Staff",
		"Mentor",
		"Applicant",
		"Attendee",
		"User",
		"Sponsor"
	]
}
```

GET /auth/roles/list/ROLE/
--------------------------

Gets the list of users with the role `ROLE`.

Request requires no body.

```json title="Example response"
{
	"userIds": [
		"google901283019238091820933",
		"google908290138109283982388",
		"github1290381"
	]
}
```
