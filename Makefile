# Makefile

GO=		go
GOBIN=  ${GOPATH}/bin

VPATH= dice:cmd/dices-go

OPTS=   -ldflags="-s -w" -v

all: dices-go

dices-go: dices.go cmds.go dice.go result.go roll.go
	${GO} build ${OPTS} -v ./cmd/...

install:
	${GO} install ${OPTS} ./cmd/...

clean:
	${GO} clean -v .

push:
	git push --all
	git push --tags

test:
	${GO} test -v ./...

lint:
	gometalinter .
