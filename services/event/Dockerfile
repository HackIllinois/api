FROM ubuntu:12.04

EXPOSE 8010

WORKDIR /opt/hackillinois/

ADD hackillinois-api-event /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x api-event

CMD ["./hackillinois-api-event"]
