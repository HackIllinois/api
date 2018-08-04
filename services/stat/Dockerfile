FROM ubuntu:12.04

EXPOSE 8011

WORKDIR /opt/hackillinois/

ADD api-stat /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x api-stat

CMD ["./api-stat"]
