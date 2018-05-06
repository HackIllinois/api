Auth
====

This is the authorization and authentication microservice supporting hackillinois. This service allows users to sign in via a specified oauth provider.

Setup
-----

First download the source:
```
go get -u github.com/HackIllinois/api-auth
```

Move into the source directory:
```
cd $GOPATH/src/github.com/HackIllinois/api-auth
```

And install the service
```
go install
```

This will place an executable in your `$GOPATH/bin` directory which you can run to start up the service.

This service uses environment variables to set the oauth proivder's client id, oauth secret, and token generation secret. Ensure that you set these environment variables before you start the service.
