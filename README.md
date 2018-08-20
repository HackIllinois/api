# HackIllinois API
This repository contains the code which runs the backend services supporting HackIllinois.

1. [Building, Testing, and Running the API](#building-testing-and-running-the-api)
2. [Release Deployment](#release-deployment)
3. [Contributing](#contributing)
4. [Documentation](#documentation)

## Building, Testing, and Running the API
In order to simply API development `make` is used for building, testing, and running the API. All `make` commands can be run from the root of the repository and they will properly find and operate on all of the services.

### Building the API
Run the following command from the root of the repository. The gateway and all services will be built to `$GOPATH/bin` which should be in your `PATH` so you can start up any of them easily.
```
make all
```

### Testing the API
Run the following command from the root of the repository. The gateway and all services will have their tests run and any errors will be reported.
```
make test
```

### Running the API
Run the following command from the root of the repository. Your `PATH` must contain `$GOPATH/bin` since that is where the services will be started from. Note that this command will not rebuild the API so you must first build the API to ensure your binaries are up to date.
```
make run
```

## Release Deployment
For release deployment see the `README.md` file in `release/`.

## Contributing
For contributing guidelines see the `CONTRIBUTING.md` file in the repository root.

## Documentation
Documentation will be available soon.
