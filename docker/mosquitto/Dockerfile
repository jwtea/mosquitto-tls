FROM ubuntu:trusty

ARG PG_HOST=127.0.0.1
ARG PG_PORT=5432

ARG PG_USER=docker
ARG PG_PASS=docker

RUN apt-get update -y && apt-get install -y wget git && \
  wget http://mosquitto.org/files/source/mosquitto-1.4.14.tar.gz && \
  apt-get -y install build-essential uuid-dev libssl-dev libpq-dev && \
  ldconfig && \
  mkdir /mosquitto && tar -xvzf mosquitto-1.4.14.tar.gz -C /mosquitto && \
  adduser --system --disabled-password --disabled-login mosquitto

COPY config/mosquitto.conf /mosquitto/mosquitto-1.4.14/mosquitto.conf
COPY certs /mosquitto/certs

RUN cd /mosquitto/mosquitto-1.4.14 && make WITH_SRV:=no
RUN mkdir /var/lib/mosquitto

EXPOSE 8883
CMD ["/mosquitto/mosquitto-1.4.14/src/mosquitto", "-c", "/mosquitto/mosquitto-1.4.14/mosquitto.conf"]