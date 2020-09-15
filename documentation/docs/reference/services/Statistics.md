Event
=====

GET /stat/
----------

Returns statistics for all services.

Response format:
```
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

Response format:
```
{

	"OpeningCeremony": 8,
	"Breakfast": 6
}
```
