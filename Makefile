# Makefile

VPATH= dice:cmd/dices-go

all: dices-go

dices-go: dices.go dice.go roll.go
	go build -v ./...
	go test -v ./...

clean:
	go clean -v

push:
	git push --all
