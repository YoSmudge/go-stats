
GLIDEBIN=${GOPATH}/bin/glide
GOFILES=$(shell find . -path ./vendor -prune -o -name '*.go' -print)

$(GLIDEBIN):
	go get github.com/Masterminds/glide

test-setup: $(GLIDEBIN)
	${GOPATH}/bin/glide install

test: test-setup
	go build
	go test -v ./
	gofmt -l ${GOFILES} | read && echo "gofmt failures" && gofmt -d ${GOFILES} && exit 1 || true
