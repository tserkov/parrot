PARROTCMD=./cmd/parrot
PARROTBIN=./bin/parrot
BUILDFLAGS=
RUNFLAGS?=-h

all: test build
build: clean bundle-app pre-build build-linux64 post-build

pre-build:
	rice embed-go -v -i ./pkg/webserver
post-build:
	upx --brute $(PARROTBIN)*
test:
	go test -v ./...
clean:
	rm -f $(PARROTBIN)*
run:
	go build -o $(PARROTBIN) -v $(PARROTCMD)
	$(PARROTBIN) $(RUNFLAGS)
deps:
	GO111MODULE=on go mod tidy

bundle-app:
	cd app && yarn && yarn build

build-linux64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(PARROTBIN)_linux_amd64 -v $(PARROTCMD)
