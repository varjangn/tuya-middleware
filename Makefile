build:
	@go build -o bin/tuya-middleware ./cmd/api/main.go

run: build
	@./bin/tuya-middleware

test:
	@go test -f ./ ...
