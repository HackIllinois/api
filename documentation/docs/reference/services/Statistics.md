Event
=====

GET /stat/
----------

Returns statistics for all services.

Request requires no body.

```json title="Example response"
{
	"registration": {
		"school": {
			"University of Illinois Urbana-Champaign": 5,
			"Northwestern University": 3
		},
		"major": {
			"Computer Science": 4,
			"Computer Engineering": 4
		}
	},
	"event": {
		"OpeningCeremony": 8,
		"Breakfast": 6
	}
}
```

GET /stat/SERVICENAME/
----------------------

Returns statistics for the service with the name `SERVICENAME`.

Request requires no body.

```json title="Example response"
{

	"OpeningCeremony": 8,
	"Breakfast": 6
}
```
