ifndef $(GOPATH)
    GOPATH=$(shell go env GOPATH)
    export GOPATH
endif

test:
	go test golangtasks/hw1
	go test golangtasks/hw2
	go test golangtasks/hw3

install:
	sudo -S curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.39.0

linter:
	golangci-lint run
