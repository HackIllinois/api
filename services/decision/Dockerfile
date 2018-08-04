FROM ubuntu:12.04

EXPOSE 8005

WORKDIR /opt/hackillinois/

ADD hackillinois-api-decision /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x api-decision

CMD ["./hackillinois-api-decision"]
