lint:
	golangci-lint run -v
test:
	go clean -testcache && go test -v -cover ./...
run:
	go run cmd/main.go
docker-build:
	docker build -t stori-transaction-summary .
docker-run:
	docker run -p 8080:8080 stori-transaction-summary
docker-compose:
	docker-compose up -d

