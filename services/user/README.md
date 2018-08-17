User
====

This is the user microservice supporting hackillinois. This service allows basic user information such as username and email to be stored.

Setup
-----

First download the source:
```
go get -u github.com/ReflectionsProjections/api/services/user
```

Move into the source directory:
```
cd $GOPATH/src/github.com/ReflectionsProjections/api/services/user
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
USER_DB_HOST
USER_DB_NAME
```
