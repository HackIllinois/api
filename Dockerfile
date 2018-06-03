FROM ubuntu:12.04

EXPOSE 8000

WORKDIR /opt/hackillinois/

ADD api-gateway /opt/hackillinois/

RUN chmod +x api-gateway
RUN mkdir log/
RUN touch log/access.log

CMD ["./api-gateway", "-u"]
