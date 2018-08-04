FROM ubuntu:12.04

EXPOSE 8003

WORKDIR /opt/hackillinois/

ADD api-user /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x api-user

CMD ["./api-user"]
