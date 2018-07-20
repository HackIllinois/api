Mail
====

POST /mail/send/
-----------

Sends an email to the users in the list `ids` with based on the given template with generated substitutions.

Request format:
```
{
	"ids": [
		"testuser1"
	],
	"template": "api-test"
}
```

Response format:
```
{
	"results": {
		"total_rejected_recipients": 0,
		"total_accepted_recipients": 1
	}
}
```

POST /mail/send/list/
----------------

Sends an email to the users in the mailing list `listId` with based on the given template with generated substitutions.

Request format:
```
{
	"listId": "test",
	"template": "api-test"
}
```

Response format:
```
{
	"results": {
		"total_rejected_recipients": 0,
		"total_accepted_recipients": 1
	}
}
```

POST /mail/list/create/
------------------

Creates a mailing list with the ID `id` and with the initial list of users in `userIds`, if provided.  

Request format:
```
{
	"id": "test",
	"userIds": [
		"testuser1"
	]
}
```

Response format:
```
{
	"id": "test",
	"userIds": [
		"testuser1"
	]
}
```

POST /mail/list/add/
---------------

Modifies the mailing list with the ID `id` adding the users in the list `userIds`.

Request format:
```
{
	"id": "test",
	"userIds": [
		"testuser2"
	]
}
```

Response format:
```
{
	"id": "test",
	"userIds": [
		"testuser1",
		"testuser2"
	]
}
```

POST /mail/list/remove/
------------------

Modifies the mailing list with the ID `id` adding the removing in the list `userIds`.

Request format:
```
{
	"id": "test",
	"userIds": [
		"testuser2"
	]
}
```

Response format:
```
{
	"id": "test",
	"userIds": [
		"testuser1"
	]
}
```

GET /mail/list/LISTID/
-----------------

Returns the mailing list with the ID `LISTID`.

Response format:
```
{
	"id": "test",
	"userIds": [
		"testuser1"
	]
}
```
