vet:
	go vet ./...

fmt:
	go fmt ./...

test:
	go test -v ./...
	test:

mod:
	go mod tidy

build:
	go build -o pomodoro ./cmd/pomodoro

run:
	go run ./cmd/pomodoro

clean:
	rm -f pomodoro

