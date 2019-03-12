Notifications
=============

GET /notifications/topic/
-------------------------

Returns the list of valid topics.

Response format:
```
{
	"topics": [
		"Admin",
		"Staff",
		"Mentor",
		"Applicant",
		"Attendee",
		"User",
		"Sponsor",
		"ExampleTopic"
	]
}
```

POST /notifications/topic/
--------------------------

Create a new topic with the specified information.

Request format:
```
{
	"id": "ExampleTopic"
}
```

Response format:
```
{
	"id": "ExampleTopic",
	"userIds": []
}
```

GET /notifications/topic/all/
-----------------------------

Returns the notifications for all topics the user is subscribed to.

Response format:
```
{
	"notifications": [
		{
			"id": "52fdfc072182654f163f5f0f9a621d72",
			"topic": "User",
			"title": "Example Title",
			"body": "Example Body",
			"time": 1551805897
		}
	]
}
```

GET /notifications/topic/public/
-----------------------------

Returns the notifications which are publically viewable.

Response format:
```
{
	"notifications": [
		{
			"id": "52fdfc072182654f163f5f0f9a621d72",
			"topic": "User",
			"title": "Example Title",
			"body": "Example Body",
			"time": 1551805897
		}
	]
}
```

GET /notifications/topic/TOPICID/
-----------------------------

Returns the notifications for the topic with the id `TOPICID`.

Response format:
```
{
	"notifications": [
		{
			"id": "52fdfc072182654f163f5f0f9a621d72",
			"topic": "User",
			"title": "Example Title",
			"body": "Example Body",
			"time": 1551805897
		}
	]
}
```

POST /notifications/topic/TOPICID/
----------------------------------

Publishes a notification to the topic with the ID `TOPICID`.

Request format:
```
{
	"title": "Example Title",
	"body": "Example Body"
}
```

Response format:
```
{
	"success": 5,
	"failure": 0
}
```

DELETE /notifications/topic/TOPICID/
------------------------------------

Deletes the topic with the ID `TOPICID`.

Response format:
```
{}
```

POST /notifications/topic/TOPICID/subscribe/
--------------------------------------------

Subscribes the user to the topic with the id `TOPICID` and return the user's list of subscriptions.

Response format:
```
{
	"topics": [
		"ExampleTopic",
		"User",
		"Applicant",
		"Admin",
		"Attendee",
		"Mentor"
	]
}
```

POST /notifications/topic/TOPICID/unsubscribe/
--------------------------------------------

Unsubscribes the user to the topic with the id `TOPICID` and return the user's list of subscriptions.

Response format:
```
{
	"topics": [
		"User",
		"Applicant",
		"Admin",
		"Attendee",
		"Mentor"
	]
}
```

POST /notifications/device/
---------------------------

Registers the specified device token to the current user.

Request format:
```
{
	"token": "example_token",
	"platform": "android"
}
```

Response format:
```
{
    "devices": [
        "arn:example139091820398"
    ]
}
```
