LOGGER_NAME = logger
NODE_NAME = node

.PHONY: all logger node run clean

all: logger node run
 
logger:
	go build -o ./bin/${LOGGER_NAME} logger/logger.go

node:
	go build -o ./bin/${NODE_NAME} node/node.go
 
run:
	./bin/${LOGGER_NAME} 1234

clean:
	go clean
	rm -f ./bin/${LOGGER_NAME} ./bin/${NODE_NAME}