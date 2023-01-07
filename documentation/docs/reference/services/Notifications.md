Notifications
=============

GET /notifications/topic/
-------------------------

Returns the list of valid topics.

Request requires no body.

```json title="Example response"
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

```json title="Example request"
{
	"id": "ExampleTopic"
}
```

```json title="Example response"
{
	"id": "ExampleTopic",
	"userIds": []
}
```

GET /notifications/topic/all/
-----------------------------

Returns the notifications for all topics the user is subscribed to.

Request requires no body.

```json title="Example response"
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

Returns the notifications which are publicly viewable.

Request requires no body.

```json title="Example response"
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

Request requires no body.

```json title="Example response"
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

Publishes a notification to the topic with the ID `TOPICID`. The `id` in the response is the ID for the notification order which is sending the actual notifications asynchronously.

```json title="Example request"
{
	"title": "Example Title",
	"body": "Example Body"
}
```

```json title="Example response"
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"recipients": 5,
	"success": 0,
	"failure": 0,
	"time": 1553564589
}
```

DELETE /notifications/topic/TOPICID/
------------------------------------

Deletes the topic with the ID `TOPICID`.

Request requires no body.

```json title="Example response"
{}
```

POST /notifications/topic/TOPICID/subscribe/
--------------------------------------------

Subscribes the user to the topic with the id `TOPICID` and return the user's list of subscriptions.

Request requires no body.

```json title="Example response"
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

Request requires no body.

```json title="Example response"
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

```json title="Example request"
{
	"token": "example_token",
	"platform": "android"
}
```

```json title="Example response"
{
	"devices": [
		"arn:example139091820398"
	]
}
```

GET /notifications/order/ID/
----------------------------

Returns the notification order with the `id` ID. This endpoint should be used to determine the status of an asynchronously published notification.

Request requires no body.

```json title="Example response"
{
	"id": "52fdfc072182654f163f5f0f9a621d72",
	"recipients": 5,
	"success": 5,
	"failure": 0,
	"time": 1553564589
}
```
