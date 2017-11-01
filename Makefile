all: build

test:
	go test -v -cover ./...

lint:
	go vet ./...
	golint ./...

install-deps:
	go get -d -t ./...
	go get github.com/golang/lint/golint

build: test lint clean
	GOOS=darwin GOARCH=amd64 go build -o build/sysrq_darwin_amd64 ./cmd/sysrq
	GOOS=linux  GOARCH=amd64 go build -o build/sysrq_linux_amd64  ./cmd/sysrq

CMD := help

trigger: build
	vagrant ssh -c 'sudo /vagrant/build/sysrq_linux_amd64 $(CMD)'

log:
	tail -F ubuntu-*-console.log

clean:
	$(RM) -r build

.PHONY: build
