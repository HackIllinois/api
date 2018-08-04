Decision
========

This is the decision microservice supporting hackillinois. This service allows decisions on attendee application to be stored and queried.

Setup
-----

First download the source:
```
go get -u github.com/HackIllinois/api-decision
```

Move into the source directory:
```
cd $GOPATH/src/github.com/HackIllinois/api-decision
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
DECISION_DB_HOST
DECISION_DB_NAME
```
