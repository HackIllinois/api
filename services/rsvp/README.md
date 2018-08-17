RSVP
====

This is the rsvp microservice supporting hackillinois. This service allows an applicant's rsvp to be created, updated, and retreived.

Setup
-----

First download the source:
```
go get -u github.com/ReflectionsProjections/api/services/rsvp
```

Move into the source directory:
```
cd $GOPATH/src/github.com/ReflectionsProjections/api/services/rsvp
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
RSVP_DB_HOST
RSVP_DB_NAME
```
