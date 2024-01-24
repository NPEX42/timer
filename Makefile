all: generate build run


generate:
	templ generate components/

build:
	go build ./app/main.go

run:
	go run ./app/main.go