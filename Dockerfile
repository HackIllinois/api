FROM ubuntu:12.04

EXPOSE 8004

WORKDIR /opt/hackillinois/

ADD api-registration /opt/hackillinois/

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN chmod +x api-registration

CMD ["./api-registration"]
