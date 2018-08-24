Print
=======

This is the print microservice supporting HackIllinois. This service allows publication of print jobs to print nodes.

Setup
-----

First download the source:
```
go get -u github.com/HackIllinois/api/services/print
```

Move into the source directory:
```
cd $GOPATH/src/github.com/HackIllinois/api/services/print
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
AWS_REGION
```
