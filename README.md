# HackIllinois API
[![Build Status](https://travis-ci.com/HackIllinois/api.svg?branch=master)](https://travis-ci.com/HackIllinois/api)
[![Go Report Card](https://goreportcard.com/badge/github.com/HackIllinois/api)](https://goreportcard.com/report/github.com/HackIllinois/api)

This repository contains the code which runs the backend services supporting HackIllinois.

1. [Developer Environment Setup](#developer-environment-setup)
2. [Building, Testing, and Running the API](#building-testing-and-running-the-api)
3. [Release Deployment](#release-deployment)
4. [Contributing](#contributing)
5. [Documentation](#documentation)

## Developer Environment Setup
In order to work on the API there are a few setups neccessary in order to setup your developer environemnt.

### Installing Dependencies
We highly reccommend that you use Ubuntu 18.04 when working on API development. The API is written and Go and makes use of MongoDB for storing data. You will have to install both of these before working on the API. You will also need a few common development tools including `make` and `git`.

#### Installing Development Tools
Both `make` and `git` can be installed from the default ubuntu package repositories. Run the following commands to install both tools. You may need to run the commands under `sudo`.
```
apt-get update
apt-get install build-essential git
```

#### Installing Go
Follow the [Go Installation Instructions](https://golang.org/doc/install#install) for installing Go. Run `go version` to ensure that Go has been properly installed.

#### Installing MongoDB
Follow the [MongoDB Installation Instructions](https://docs.mongodb.com/manual/installation/#mongodb-community-edition) for installing MongoDB. Once MongoDB is installed ensure `mongod` is running. If it is not running then start the service.

#### Downloading the API source and dependencies
Run the following command to retrieve the API and all it's dependencies. The API source will be cloned to a folder called `api` in your current directory.
```
git clone https://github.com/HackIllinois/api.git
```

#### First time API setup
After downloading the API source code you will want to build the entire repository and generate an `Admin` token for testing endpoints. This can be done by moving into the API folder and running:
```
make setup
```
You should see your `Admin` token near the bottom of the output.

#### Useful tools for development
There are a couple other useful but not necessary tools for working on the API. The first is a GUI tool for viewing and modifying the database. There are many options including [MongoDB Compass](https://www.mongodb.com/products/compass) and [Robo 3T](https://robomongo.org/). You will also want to install [Postman](https://www.getpostman.com/) for making requests to the API when testing.

## Building, Testing, and Running the API
In order to simply API development `make` is used for building, testing, and running the API. All `make` commands can be run from the root of the repository and they will properly find and operate on all of the services.

### Building the API
Run the following command from the root of the repository. The gateway and all services will be built into `bin`.
```
make all
```

### Testing the API
Run the following command from the root of the repository. The gateway and all services will have their tests run and any errors will be reported.
```
make test
```

### Running the API
Run the following command from the root of the repository. Note that this command will not rebuild the API so you must first build the API to ensure your binaries are up to date.
```
make run
```

## API Container
For instructions on building the API container see the `README.md` file in `container/`.

## Contributing
For contributing guidelines see the `CONTRIBUTING.md` file in the repository root.

## Documentation
We use [MkDocs](https://www.mkdocs.org/) for our documentation, and host at [HackIllinois Docs](https://docs.hackillinois.org/).
