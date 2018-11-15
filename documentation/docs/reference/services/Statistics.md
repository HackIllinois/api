Event
=====

GET /stat/service/SERVICENAME/
------------------------------

Returns the service with the name of `SERVICENAME`.

Response format:
```
{
	"name": "ExampleService",
	"url": "http://localhost:8050"
}
```

POST /stat/service/
-------------------

Registers a service with the given name. Returns the registered service.

Request format:
```
{
	"name": "ExampleService",
	"url": "http://localhost:8050"
}
```

Response format:
```
{
	"name": "ExampleService",
	"url": "http://localhost:8050"
}
```
