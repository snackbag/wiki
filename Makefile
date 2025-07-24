SOURCES := *.go

.PHONY:	clean

all: wiki

clean:
	rm -rf wiki

wiki: Makefile $(SOURCES)
	go mod tidy
	go build

dev:
	go mod tidy
	go run .
