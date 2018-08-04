FROM ubuntu:12.04

EXPOSE 8009

WORKDIR /opt/hackillinois/

ADD api-mail /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x api-mail

CMD ["./api-mail"]
