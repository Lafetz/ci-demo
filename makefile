test:
	go test  ./app 
coverage:
	mkdir -p coverage
	go test  -coverprofile=coverage.out ./app/... ;
	go tool cover -func=coverage.out -o coverage/coverage.txt
lint:
	golangci-lint run
build:
	go build -ldflags="-s" -o ./bin/web ./