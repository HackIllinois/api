FROM ubuntu:18.04

# Expose gateway port options
EXPOSE 8000
EXPOSE 80
EXPOSE 443

# All hackillinois software will be located here
WORKDIR /opt/hackillinois/

# Add gateway and services to container
ADD hackillinois-api /opt/hackillinois/

# Add statup script to container
ADD start.sh /opt/hackillinois/

# Install certificates for https
RUN apt-get update
RUN apt-get install -y ca-certificates

# Make gateway and services executable
RUN chmod +x hackillinois-api

# Make the startup script executable
RUN chmod +x start.sh

# Create the logging directory for the gateway
RUN mkdir log/
RUN touch log/access.log

# Startup gateway and services
CMD ["./start.sh"]
