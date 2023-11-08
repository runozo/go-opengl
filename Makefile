.PHONY: build

build:
	go build -ldflags="-s -w -v" -o ./bin/demo ./src
#	upx -9 ./bin/demo

.PHONY: run

run: build
	./bin/demo
