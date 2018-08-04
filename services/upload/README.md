Upload
======

This is the upload microservice supporting hackillinois. This service allows users to upload files for storage and retreival by staff.

Setup
-----

First download the source:
```
go get -u github.com/HackIllinois/api/services/upload
```

Move into the source directory:
```
cd $GOPATH/src/github.com/HackIllinois/api/services/upload
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
UPLOAD_DB_HOST
UPLOAD_DB_NAME
```
