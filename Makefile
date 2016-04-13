
GOFILES=$(shell find . -path ./vendor -prune -o -name '*.go' -print)

test:
	go build
	go test -v ./
	gofmt -l ${GOFILES} | read && echo "gofmt failures" && gofmt -d ${GOFILES} && exit 1 || true
