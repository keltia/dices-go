# Makefile

VPATH= dice

all: dices-go

dices-go: dices.go dice.go roll.go
	go build -v -o dices-go
	go test -v ./...

clean:
	go clean -v

push:
	git push --all
