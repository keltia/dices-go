# Makefile

VPATH= dice

all: dices-go

dices: dices.go dice.go roll.go
	go build -v -o dices-go

clean:
	go clean -v

push:
	git push --all
