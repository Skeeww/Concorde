run:
	go run main.go config.go queue.go protocol.go http.go osc.go websocket.go random.go

build:
	go build -o ./build/concorde main.go config.go queue.go protocol.go http.go osc.go websocket.go random.go
