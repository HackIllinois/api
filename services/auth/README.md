Auth
====

This is the authorization and authentication microservice supporting hackillinois. This service allows users to sign in via a specified oauth provider.

Setup
-----

First download the source:
```
go get -u github.com/ReflectionsProjections/api/services/auth
```

Move into the source directory:
```
cd $GOPATH/src/github.com/ReflectionsProjections/api/services/auth
```

And install the service
```
go install
```

This will place an executable in your `$GOPATH/bin` directory which you can run to start up the service.

Environment Variables
---------------------

This service uses environment variables to set the oauth proivder's client id, oauth secret, and token generation secret. Ensure that you set these environment variables before you start the service.

Current environment variables used:
```
GITHUB_CLIENT_ID
GITHUB_CLIENT_SECRET
TOKEN_SECRET
AUTH_DB_HOST
AUTH_DB_NAME
AUTH_PORT
USER_SERVICE
```
