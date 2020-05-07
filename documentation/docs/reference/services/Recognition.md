Recognition
=====


GET /recognition/
---------------------

Returns a list of all recognitions.

Response format:
```
{
    "recognitions": [
        {
            "id": "81855ad8681d0d86d1e91e00167939cb",
            "name": "Example Recognition 10",
            "description": "This is a description",
            "presenter": "Example presenter",
            "recognitionId": "81855ad8681d0d86d1e91e00167939cb",
            "recipients": [
                {
                    "type": "PROJECT",
                    "typeId": "52fdfc072182654f163f5f0f9a621d72"
                }
            ],
            "tags": [
                "Data Science",
                "Mobile"
            ]
        },
        {
            "id": "9566c74d10037c4d7bbb0407d1e2c649",
            "name": "Best Computer Security",
            "description": "Good mastery of safe coding practices",
            "presenter": "HackIllinois",
            "recognitionId": "52fdfc072182654f163f5f0f9a621d72",
            "recipients": [
                {
                    "type": "INDIVIDUAL",
                    "typeId": "github09829234"
                }
            ],
            "tags": [
                "Security",
                "Virus"
            ]
        }
    ]
}
```

POST /recognition/
-----------

Creates an recognition with the requested fields. Returns the created recognition.

Request format:
```
{
	"name": "Example Recognition 10",
	"description": "This is a description",
	"presenter": "Example presenter",
	"recognitionId": "81855ad8681d0d86d1e91e00167939cb",
	"tags": ["Data Science", "Mobile"],
	"recipients": [
		{
			"type": "ALL"
		}
	]
}
```

Response format:
```
{
	"name": "Example Recognition 10",
	"description": "This is a description",
	"presenter": "Example presenter",
	"recognitionId": "81855ad8681d0d86d1e91e00167939cb",
	"tags": ["Data Science", "Mobile"],
	"recipients": [
		{
			"type": "ALL"
		}
	]
}
```

DELETE /recognition/RECOGNITIONID/
-----------

Endpoint to delete an recognition with name `RECOGNITIONID`

Response format:
```
{
	"name": "Example Recognition 10",
	"description": "This is a description",
	"presenter": "Example presenter",
	"recognitionId": "81855ad8681d0d86d1e91e00167939cb",
	"tags": ["Data Science", "Mobile"],
	"recipients": [
		{
			"type": "ALL"
		}
	]
}
```


GET /recognition/filter/?key=value
---------------------

Returns all recognitions, filtered with the given key-value pairs.

Response format:
```
{
    "recognitions": [
        {
            "id": "81855ad8681d0d86d1e91e00167939cb",
            "name": "Example Recognition 10",
            "description": "This is a description",
            "presenter": "Example presenter",
            "recognitionId": "81855ad8681d0d86d1e91e00167939cb",
            "recipients": [
                {
                    "type": "PROJECT",
                    "typeId": "52fdfc072182654f163f5f0f9a621d72"
                }
            ],
            "tags": [
                "Data Science",
                "Mobile"
            ]
        }
    ]
}
```
