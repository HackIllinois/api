# API Container

This API container serves two purposes:

1. Running a single service in a cluster deployment
2. Running all services in a single container deployment

Documentation for those wishing to run the API for their events is provided below.

1. [Building the API Container](#building-the-api-container)
2. [Building the API Container Image](#building-the-api-container-image)
3. [Running the API from GitHub Releases Image](#running-the-api-from-github-releases-image)
4. [Running the API from DockerHub Image](#running-the-api-from-dockerhub-image)

## Building the API Container

Building a container requires that docker and go have already been installed. The following command should be run from the root of the api repository.

```
make container
```

This will build a container for the API labelled `hackillinois-api:latest`.

## Building the API Container Image

Building a container image requires that docker and go have already been installed. The following command should be run from the root of the api repository.

```
make release
```

This will generate an archive `hackillinois-api-image.tar` in the `bin` directory containing the docker image.

Running the API from GitHub Releases Image
------------------------------------------
First install docker on your deployment target. Once ready go to [GitHub Releases](https://github.com/HackIllinois/api/releases) and download the lastest `hackillinois-api-image.tar` image archive.

Setup external dependencies:
- MongoDB Server
- S3 Bucket
- Sparkpost
- Oauth Applications
	- Github
	- Google
	- Linkedin

Setup your configuration by copying `env.template` to `env.config`. You must define all the environment variables which are in the file for the API to be fully functional. Most of these environment variables define how to access the external depencies you setup earlier.

Load the container from the archive
```
docker load -i hackillinois-api-image.tar
```

Run the container with your environment file
```
docker run -d --env-file env.config hackillinois-api:release
```

This will run all services in a single container and the API should now be accessable on port 8000.

If you would like to run a single container per service, as you would do in a cluster deployment, run the container as follows
```
docker run -d --env-file env.config hackillinois-api:release <service-name>
```

Running the API from DockerHub Image
------------------------------------
In the future we will be uploading this container to DockerHub to simplify deployment. However for now, you must follow the [Running the API from GitHub Releases Image](#running-the-api-from-github-releases-image) instructions.
