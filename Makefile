APP_NAME=server

build:
	go build -o $(APP_NAME) cmd/server/main.go

run:
	go run ./cmd/$(APP_NAME)/main.go

clean:
	rm -f $(APP_NAME)