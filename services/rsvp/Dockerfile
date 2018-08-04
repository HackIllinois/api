FROM ubuntu:12.04

EXPOSE 8006

WORKDIR /opt/hackillinois/

ADD hackillinois-api-rsvp /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x hackillinois-api-rsvp

CMD ["./hackillinois-api-rsvp"]
