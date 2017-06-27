# Makefile

VPATH= dice:cmd/dices-go
OPTS= -ldflags="-s -w" -v

all: dices-go

dices-go: dices.go dice.go roll.go
	go build ${OPTS} ./...
	go test ${OPTS} ./...

clean:
	go clean -v

push:
	git push --all
