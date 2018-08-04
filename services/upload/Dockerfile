FROM ubuntu:12.04

EXPOSE 8008

WORKDIR /opt/hackillinois/

ADD api-upload /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x api-upload

CMD ["./api-upload"]
