FROM ubuntu:12.04

EXPOSE 8008

WORKDIR /opt/hackillinois/

ADD hackillinois-api-notifications /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x hackillinois-api-notifications

CMD ["./hackillinois-api-notifications"]
