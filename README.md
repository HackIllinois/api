Event
=====

This is the event microservice supporting hackillinois. This service allows an event to be created, updated, and retreived. It also supports tracker which users attend each event.

Setup
-----

First download the source:
```
go get -u github.com/HackIllinois/api-event
```

Move into the source directory:
```
cd $GOPATH/src/github.com/HackIllinois/api-event
```

And install the service
```
go install
```

This will place an executable in your `$GOPATH/bin` directory which you can run to start up the service.

Environment Variables
---------------------

This service uses environment variables to set the mongo database used. Ensure that you set these environment variables before you start the service.

Current environment variables used:
```
EVENT_DB_HOST
EVENT_DB_NAME
```
