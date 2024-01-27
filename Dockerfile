FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go get -d -v ./...

COPY *.go ./

RUN go build -o lol-player-finder .

EXPOSE 5000

CMD [ "./lol-player-finder" ]