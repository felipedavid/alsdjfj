run: build
	@./bin/app

build: templ
	@go build -o bin/app cmd/main.go

templ:
	@templ generate
