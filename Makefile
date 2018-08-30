SOURCES = main.go

all : gotcpserver

gotcpserver: $(SOURCES)
	go build -x -o gotcpserver $(SOURCES)

clean:
	go clean -x
	rm -vf gotcpserver

check:
	go test -v .

install:
	go install -v .

.PHONY: all clean check install
