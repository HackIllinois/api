FROM ubuntu:18.04

# Expose gateway port options
EXPOSE 8000
EXPOSE 80
EXPOSE 443

# All hackillinois software will be located here
WORKDIR /opt/hackillinois/

# Add gateway and services to container
ADD hackillinois-api-* /opt/hackillinois/

# Add statup script to container
ADD start.sh /opt/hackillinois/

# Install certificates for https
RUN apt-get update
RUN apt-get install -y ca-certificates

# Make gateway and services executable
RUN chmod +x hackillinois-api-*

# Make the startup script executable
RUN chmod +x start.sh

# Create the logging directory for the gateway
RUN mkdir log/
RUN touch log/access.log

# Setup default environment variables
ENV IS_PRODUCTION true

ENV AUTH_SERVICE http://localhost:8002
ENV USER_SERVICE http://localhost:8003
ENV REGISTRATION_SERVICE http://localhost:8004
ENV DECISION_SERVICE http://localhost:8005
ENV RSVP_SERVICE http://localhost:8006
ENV CHECKIN_SERVICE http://localhost:8007
ENV UPLOAD_SERVICE http://localhost:8008
ENV MAIL_SERVICE http://localhost:8009
ENV EVENT_SERVICE http://localhost:8010
ENV STAT_SERVICE http://localhost:8011

ENV AUTH_PORT :8002
ENV USER_PORT :8003
ENV REGISTRATION_PORT :8004
ENV DECISION_PORT :8005
ENV RSVP_PORT :8006
ENV CHECKIN_PORT :8007
ENV UPLOAD_PORT :8008
ENV MAIL_PORT :8009
ENV EVENT_PORT :8010
ENV STAT_PORT :8011

ENV GATEWAY_PORT 8000

ENV AUTH_DB_NAME auth
ENV USER_DB_NAME user
ENV REGISTRATION_DB_NAME registration
ENV DECISION_DB_NAME decision
ENV RSVP_DB_NAME rsvp
ENV CHECKIN_DB_NAME checkin
ENV MAIL_DB_NAME mail
ENV EVENT_DB_NAME event
ENV STAT_DB_NAME stat

ENV SPARKPOST_API https://api.sparkpost.com/api/v1

# Startup gateway and services
CMD ["./start.sh"]
