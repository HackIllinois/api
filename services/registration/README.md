Registration
============

This is the registration microservice supporting hackillinois. This service allows user registrations to be created, updated, and retreived.

Setup
-----

First download the source:
```
go get -u github.com/ReflectionsProjections/api/services/registration
```

Move into the source directory:
```
cd $GOPATH/src/github.com/ReflectionsProjections/api/services/registration
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
REGISTRATION_DB_HOST
REGISTRATION_DB_NAME
```
