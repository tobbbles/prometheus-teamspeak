FROM golang:1.6.3-alpine

RUN \
  apk add --update git \
  && rm -rf /var/cache/apk/*

RUN go get github.com/constabulary/gb/...

WORKDIR /app

RUN git clone https://github.com/mnzt/prometheus-teamspeak.git /app \
  && cd /app \
  && gb build

COPY /app/bin/prometheus-teamspeak /usr/local/bin

ENTRYPOINT ["prometheus-teamspeak"]

EXPOSE 8010
