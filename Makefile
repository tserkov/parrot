PARROTCMD=./cmd/parrot
PARROTBIN=./bin/parrot
BUILDFLAGS=
RUNFLAGS?=-h

all: test build
build: clean pre-build build-linux64 post-build

pre-build:
	packr
post-build:
	upx --brute $(PARROTBIN)*
test:
	go test -v ./...
clean:
	go clean
	rm -f $(PARROTBIN)*
run:
	go build -o $(PARROTBIN) -v $(PARROTCMD)
	$(PARROTBIN) $(RUNFLAGS)
deps:
	GO111MODULE=on go mod tidy

build-linux64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(PARROTBIN)_linux_amd64 -v $(PARROTCMD)
