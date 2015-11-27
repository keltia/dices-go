# Makefile

VPATH= dice

all: dices

dices: dices.go dice.go roll.go
	go build -v -o dices

clean:
	go clean -v

push:
	git push --all
