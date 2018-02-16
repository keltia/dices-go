# Makefile

GOBIN=   ${GOPATH}/bin

VPATH= dice:cmd/dices-go

OPTS=   -ldflags="-s -w" -v

all: dices-go

dices-go: dices.go cmds.go dice.go roll.go
	go build ${OPTS} -v ./cmd/...
	go test -v ./...

install:
	go install ${OPTS} ./cmd/...

clean:
	go clean -v .

push:
	git push --all

test:
	go test -v ./...

lint:
	gometalinter .
