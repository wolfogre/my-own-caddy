VERSION=$(shell go run main.go -version)

tidy:
	go mod tidy -v
build:
	rm -rf dist
	go run main.go
	./readme.sh
	git diff HEAD --quiet || exit 1
publish: build
	hub release create -a dist/caddy_$(VERSION)_darwin_amd64 -m v$(VERSION) v$(VERSION)
