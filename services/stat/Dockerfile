FROM ubuntu:12.04

EXPOSE 8011

WORKDIR /opt/hackillinois/

ADD hackillinois-api-stat /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x hackillinois-api-stat

CMD ["./hackillinois-api-stat"]
