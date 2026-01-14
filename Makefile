BINARY_NAME=stickbug
MAIN_PATH=./app/main.go

.PHONY: all linux windows clean

all: linux windows

linux:
	go build -v -o bin/$(BINARY_NAME) $(MAIN_PATH)

windows:
	CGO_ENABLED=1 \
	GOOS=windows \
	GOARCH=amd64 \
	CC=x86_64-w64-mingw32-gcc \
	go build -v -o bin/$(BINARY_NAME).exe $(MAIN_PATH)

clean:
	rm -rf bin/
