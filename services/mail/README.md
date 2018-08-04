Mail
====

This is the mail microservice supporting hackillinois. This service allows templatized mail to be sent to user by id.

Setup
-----

First download the source:
```
go get -u github.com/HackIllinois/api-mail
```

Move into the source directory:
```
cd $GOPATH/src/github.com/HackIllinois/api-mail
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
MAIL_DB_HOST
MAIL_DB_NAME
```
