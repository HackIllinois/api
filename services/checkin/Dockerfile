FROM ubuntu:12.04

EXPOSE 8007

WORKDIR /opt/hackillinois/

ADD hackillinois-api-checkin /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x api-checkin

CMD ["./hackillinois-api-checkin"]
