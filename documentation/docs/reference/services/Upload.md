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
