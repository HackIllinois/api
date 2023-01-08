Mail
====

POST /mail/send/
-----------

Sends an email to the users in the list `ids` with based on the given template with generated substitutions.

```json title="Example request"
{
	"ids": [
		"testuser1"
	],
	"template": "api-test"
}
```

```json title="Example response"
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

```json title="Example request"
{
	"listId": "test",
	"template": "api-test"
}
```

```json title="Example response"
{
	"results": {
		"total_rejected_recipients": 0,
		"total_accepted_recipients": 1
	}
}
```

GET /mail/list/
---------------------

Returns a list of all created mailing lists.

Request requires no body.

```json title="Example response"
{
	"mailLists": [
		{
			"id": "test",
			"userIds": [
				"testuser1"
			]
		},
		{
			"id": "test2",
			"userIds": [
				"testuser2",
				"testuser3"
			]
		}
	]
}
```

POST /mail/list/create/
------------------

Creates a mailing list with the ID `id` and with the initial list of users in `userIds`, if provided.  

```json title="Example request"
{
	"id": "test",
	"userIds": [
		"testuser1"
	]
}
```

```json title="Example response"
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

```json title="Example request"
{
	"id": "test",
	"userIds": [
		"testuser2"
	]
}
```

```json title="Example response"
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

Modifies the mailing list with the ID `id`, removing users that are specified in the list `userIds`.

```json title="Example request"
{
	"id": "test",
	"userIds": [
		"testuser2"
	]
}
```

```json title="Example response"
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

Request requires no body.

```json title="Example response"
{
	"id": "test",
	"userIds": [
		"testuser1"
	]
}
```
