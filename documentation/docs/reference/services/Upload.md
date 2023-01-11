Upload
======

GET /upload/resume/USERID/
--------------------------

Returns the resume stored for the user with the `id` `USERID`.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"resume": "https://bucket.s3.amazonaws.com/resume.pdf"
}
```

GET /upload/resume/
-------------------

Returns the resume stored for the currently authenticated user (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"resume": "https://bucket.s3.amazonaws.com/resume.pdf"
}
```

GET /upload/resume/upload/
--------------------------

Returns the S3 link for resume uploading for the currently authenticated user (determined by the JWT in the `Authorization` header). The user's resume can be `PUT` to this link.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"resume": "https://bucket.s3.amazonaws.com/resume.pdf"
}
```

GET /upload/photo/USERID/
--------------------------

Returns the photo stored for the user with the `id` `USERID`.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"photo": "https://bucket.s3.amazonaws.com/photo"
}
```

GET /upload/photo/
-------------------

Returns the photo stored for the currently authenticated user (determined by the JWT in the `Authorization` header).

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"photo": "https://bucket.s3.amazonaws.com/photo"
}
```

GET /upload/photo/upload/
--------------------------

Returns the S3 link for photo uploading for the currently authenticated user (determined by the JWT in the `Authorization` header). The user's photo can be `PUT` to this link.

Request requires no body.

```json title="Example response"
{
	"id": "github0000001",
	"photo": "https://bucket.s3.amazonaws.com/photo"
}
```

GET /upload/blobstore/ID/
-------------------------

Returns the blob stored with the `id` `ID`.

Request requires no body.

```json title="Example response"
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```

POST /upload/blobstore/
-----------------------

Creates and stores a blob with the specified `id` and `data`. `data` can be a single json field or an json object.

```json title="Example request"
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```

```json title="Example response"
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```

PUT /upload/blobstore/
----------------------

Updates the blob with the specified `id`. `data` can be a single json field or an json object.

```json title="Example request"
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```

```json title="Example response"
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```

DELETE /upload/blobstore/ID/
----------------------

Deletes the blob with the specified `id`.

```json title="Example request"
{
	"id": "exampleblob",
}
```

```json title="Example response"
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```
