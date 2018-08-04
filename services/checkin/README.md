Checkin
=======

This is the checkin microservice supporting hackillinois. This service allows checkins to checkin and tracks if they picked up swag.

Setup
-----

First download the source:
```
go get -u github.com/HackIllinois/api-checkin
```

Move into the source directory:
```
cd $GOPATH/src/github.com/HackIllinois/api-checkin
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
CHECKIN_DB_HOST
CHECKIN_DB_NAME
```
