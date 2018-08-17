API Release
===========

The purpose of providing a release image is to snapshot the state of the api and to provide a simple deployment method for events which do not require continuous deployment.

Documentation for those wishing to run the API for their events is provided below.

1. [Building a Release Image](#building-a-release-image)
2. [Running the API from GitHub Releases Image](#running-the-api-from-github-releases-image)
3. [Running the API from DockerHub Image](#running-the-api-from-dockerhub-image)

Building a Release Image
------------------------
Building a release image requires that docker and go have already been installed. The following command should be run from the root of the api repository.

```
make release
```

This will generate an archive `hackillinois-api-image.tar` in the release directory containing the docker image.

Note that the build script will cleanup after completing so no intermediates or build folder will remain.

Running the API from GitHub Releases Image
------------------------------------------
First install docker on your deployment target. Once ready go to [GitHub Releases](https://github.com/ReflectionsProjections/api/releases) and download the lastest `hackillinois-api-image.tar` image archive.

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
docker run -d -p 8000:8000 --env-file env.config hackillinois-api:release
```

The HackIllinois API should now be accessable on port 8000.

Running the API from DockerHub Image
------------------------------------
In the future we will be uploading this container to DockerHub to simplify deployment. However for now, you must follow the [Running the API from GitHub Releases Image](#running-the-api-from-github-releases-image) instructions.
