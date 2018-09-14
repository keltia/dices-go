# Makefile

GO=		go
GOBIN=  ${GOPATH}/bin

VPATH= dice:cmd/dices-go

OPTS=   -ldflags="-s -w" -v

all: dices-go

dices-go: dices.go cmds.go dice.go result.go roll.go
	${GO} build ${OPTS} -v ./cmd/...
	${GO} test -v ./...

install:
	${GO} install ${OPTS} ./cmd/...

clean:
	${GO} clean -v .

push:
	git push --all

test:
	${GO} test -v ./...

lint:
	gometalinter .
