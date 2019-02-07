# Notifications

## GET /notifications/

Returns a list of all notifications.

Response format:

```
{
	"topics": [
		"Mentors",
		"Attendees"
	]
}
```

## POST /notifications/

Creates a new topic with the requested name. Returns the created topic.

Request format:

```
{
	"name": "Mentors"
}
```

Response format:

```
{
	"name": "Mentors"
}
```

## GET /notifications/all/

Returns a list of all past notifications.

Response format:

```
{
	"notifications": [
		{
			"title": "Notification 1",
			"body": "This is a notification!",
			"time": 1541037801,
			"topicName": "Attendee"
		},
		{
			"title": "Notification 2",
			"body": "This is another notification!",
			"time": 1541069201,
			"topicName": "Attendee"
		},
		{
			"title": "Notification 3",
			"body": "This is another notification, for another topic!",
			"time": 1541169201,
			"topicName": "Mentor"
		}
	]
}
```

## GET /notifications/TOPICNAME/

Returns a list of all past notifications for a given topic `TOPICNAME`.

Response format:

```
{
	"notifications": [
		{
			"title": "Notification 1",
			"body": "This is a notification!",
			"time": 1541037801,
			"topicName": "Attendee"
		},
	{
			"title": "Notification 2",
			"body": "This is another notification!",
			"time": 1541069201,
			"topicName": "Attendee"
		}
	]
}
```

## DELETE /notifications/TOPICNAME/

Delete a topic with name `TOPICNAME`. Returns a list of all remaining topics.

Response format:

```
{
	"topics": [
		"Mentors",
		"Attendees"
	]
}
```

## POST /notifications/TOPICNAME/

Publishes and distributes a notification to all users subscribed to the topic `TOPICNAME`. Returns the created notification.

Request format:

```
{
	"title": "Message topic",
	"body": "Message to send to users"
}
```

Response format:

```
{
	"title": "Message topic",
	"body": "Message to send to users",
	"time": 1541644690,
	"topicName": "Attendee"
}
```

## GET /notifications/TOPICNAME/info/

Gets information associated by the topic `TOPICNAME`.

Response format:

```
{
	"name": "Mentors",
	"userIds": [
		"testuser1"
	]
}
```

## POST /notifications/TOPICNAME/add/

Modifies the topic `TOPICNAME`, subscribing the users in the list `userIds`.

Request format:

```
{
	"userIds": [
		"testuser1",
		"testuser2"
	]
}
```

Response format:

```
{
	"name": "Mentors",
	"userIds": [
		"testuser1"
	]
}
```

## POST /notifications/TOPICNAME/remove/

Modifies the topic `TOPICNAME`, unsubscribing the users in the list `userIds`.

Request format:

```
{
	"userIds": [
		"testuser1",
	]
}
```

Response format:

```
{
	"name": "Mentors",
	"userIds": [
		"testuser2"
	]
}
```

## POST /notifications/device/

Associates the device specified by the provided device token with the current user.
Valid platforms are `android` and `ios`.

Request format:

```
{
	"deviceToken": "abcdef",
	"platform": "android"
}
```

Response format:

```
{
	"deviceToken": "abcdef",
	"platform": "android"
}
```
