run:
	@cd cmd/app/ && go run .

# if you have air installed for hot-reload
air:
	@cd cmd/app/ && $(shell go env GOPATH)/bin/air

build:
	@echo "building the app"
	@go build ./cmd/app/main.go