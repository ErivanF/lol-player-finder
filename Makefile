build:
	@go build -o bin/lol-player-finder

run: build
	@./bin/lol-player-finder

test:
	@go test -v ./...