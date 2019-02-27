FROM ubuntu:18.04

# Expose gateway port options
EXPOSE 8000-8050

# All hackillinois software will be located here
WORKDIR /opt/hackillinois/

# Add api binary to container
ADD hackillinois-api /opt/hackillinois/

# Install certificates for https
RUN apt-get update
RUN apt-get install -y ca-certificates

# Make api executable
RUN chmod +x hackillinois-api

# Create the logging directory for the gateway
RUN mkdir log/
RUN touch log/access.log

# Startup the api
ENTRYPOINT ["./hackillinois-api", "--service"]

# Override when starting the container to run a service
CMD []
