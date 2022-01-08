.PHONY: build install

build: 
	go build .

install:
	go build . && mv github-review-requests /usr/local/bin/grr
