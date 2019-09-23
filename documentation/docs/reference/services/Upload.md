Upload
======

GET /upload/resume/USERID/
--------------------------

Returns the resume stored for the user with the `id` `USERID`.

Response format:
```
{
	"id": "github0000001",
	"resume": "https://bucket.s3.amazonaws.com/resume.pdf"
}
```

GET /upload/resume/
-------------------

Returns the resume stored for the user with the `id` stored in the given JWT in the Authorization header.

Response format:
```
{
	"id": "github0000001",
	"resume": "https://bucket.s3.amazonaws.com/resume.pdf"
}
```

GET /upload/resume/upload/
--------------------------

Returns the S3 link for resume uploading for the user with the `id` stored in the given JWT in the Authorization header. The user's resume can be `PUT` to this link.

Response format:
```
{
	"id": "github0000001",
	"resume": "https://bucket.s3.amazonaws.com/resume.pdf"
}
```

GET /upload/blobstore/ID/
-------------------------

Returns the blob stored with the `id` `ID`.

Response format:
```
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

Request format:
```
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```

Response format:
```
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

Request format:
```
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```

Response format:
```
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```

DELETE /upload/blobstore/ID/
-----------

Endpoint to delete the blob with name `ID`. It removes the `id` `ID` from the event trackers, and every user's tracker.

Response format:
```
{
	"id": "exampleblob",
	"data": {
		"thing1": "hi",
		"thing2": "hello"
	}
}
```
