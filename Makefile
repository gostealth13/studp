deps:
	go get -v

server: deps
	go run main.go

client: server
