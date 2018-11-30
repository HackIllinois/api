FROM ubuntu:12.04

EXPOSE 8000

WORKDIR /opt/hackillinois/

ADD hackillinois-api-gateway /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x hackillinois-api-gateway
RUN mkdir log/
RUN touch log/access.log

CMD ["./hackillinois-api-gateway", "-u"]
