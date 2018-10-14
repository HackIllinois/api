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

PUT /upload/resume/
-------------------

Updated the checkin for the user with the `id` stored in the given JWT in the Authorization header.

Request format:
```
The request body should contain the raw bytes of the resume pdf
```

Response format:
```
{
	"id": "github0000001",
	"resume": "https://bucket.s3.amazonaws.com/resume.pdf"
}
```
