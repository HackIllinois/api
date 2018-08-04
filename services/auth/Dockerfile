FROM ubuntu:12.04

EXPOSE 8002

WORKDIR /opt/hackillinois/

ADD hackillinois-api-auth /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x hackillinois-api-auth

CMD ["./hackillinois-api-auth"]
