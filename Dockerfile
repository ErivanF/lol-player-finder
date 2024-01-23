FROM golang:1.21

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/lol-player-finder
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 3000

CMD [ "go", "run"]