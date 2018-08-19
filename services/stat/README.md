Stat
====

This is the stat microservice supporting hackillinois. This service allows statistics to be aggregated from services.

Setup
-----

First download the source:
```
go get -u github.com/ReflectionsProjections/api/services/stat
```

Move into the source directory:
```
cd $GOPATH/src/github.com/ReflectionsProjections/api/services/stat
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
STAT_DB_HOST
STAT_DB_NAME
```
