# Makefile

all: dices

dices: dices.go
	go build -v -o dices

clean:
	go clean -v

push:
	git push --all
