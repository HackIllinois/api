Notifications
======

GET /notifications/
---------------------

Returns a list of all notification topics.

Response format:
```
{
    "topics": [
        {
            "arn": "arn:aws:sns:us-east-2:256758753660:Mentor",
            "name": "Mentors"
        },
        {
            "arn": "arn:aws:sns:us-east-2:256758753660:Attendee",
            "name": "Attendees"
        }
    ]
}
```

POST /notifications/
-----------

Creates a new topic with the requested name. Returns the created topic's name
and Amazon Resource Name (ARN).

Request format:
```
{
	"name": "Mentors"
}
```

Response format:
```
{
	"arn": "arn:aws:sns:us-east-2:256758753660:Mentors",
	"name": "Mentors"
}
```

GET /notifications/all/
---------------------

Returns a list of all past notifications.

Response format:
```
{
    "notifications": [
        {
            "message": "This is a notification!",
            "time": 1541037801,
            "topicName": "Attendee"
        },
		{
            "message": "This is another notification!",
            "time": 1541069201,
            "topicName": "Attendee"
        },
		{
            "message": "This is another notification, for another topic!",
            "time": 1541169201,
            "topicName": "Mentor"
        }
    ]
}
```

GET /notifications/TOPICNAME/
---------------------

Returns a list of all past notifications for a given topic `TOPICNAME`.

Response format:
```
{
    "notifications": [
        {
            "message": "This is a notification!",
            "time": 1541037801,
            "topicName": "Attendee"
        },
		{
            "message": "This is another notification!",
            "time": 1541069201,
            "topicName": "Attendee"
        }
    ]
}
```

DELETE /notifications/TOPICNAME/
---------------------

Delete a topic with name `TOPICNAME`. Returns a list of all remaining topics.

Response format:
```
{
    "topics": [
        {
            "arn": "arn:aws:sns:us-east-2:256758753660:Mentor",
            "name": "Mentors"
        },
        {
            "arn": "arn:aws:sns:us-east-2:256758753660:Attendee",
            "name": "Attendees"
        }
    ]
}
```

POST /notifications/TOPICNAME/
---------------------

Publishes and distributes a notification to all users subscribed to the topic `TOPICNAME`. Returns the created notification.

Request format:
```
{
	"message": "Message to send to users"
}
```

Response format:
```
{
	"message": "Message to send to users",
	"time": 1541644690,
	"topicName": "Attendee"
}
```

GET /notifications/TOPICNAME/info/
---------------------

Gets information associated by the topic `TOPICNAME`, including its ARN.

Response format:
```
{
	"arn": "arn:aws:sns:us-east-2:256758753660:Mentors",
	"name": "Mentors"
}
```
