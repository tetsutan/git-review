
NAME:=	git-review-core

default:	build

build:
	go build -o "bin/${NAME}" ./cmd/git-review

clean:
	rm -rf bin/${NAME}

install: build
	cp ./wrapper/git-review /usr/local/bin/git-review
	cp ./bin/git-review-core /usr/local/bin/git-review-core
	chmod +x /usr/local/bin/git-review
	chmod +x /usr/local/bin/git-review-core

