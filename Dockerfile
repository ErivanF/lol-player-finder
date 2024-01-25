FROM golang:1.21

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go get -u github.com/cosmtrek/air

COPY *.go ./

RUN go build -o lol-player-finder .

EXPOSE 5000

ENTRYPOINT ["air"]

CMD [ "./lol-player-finder" ]
