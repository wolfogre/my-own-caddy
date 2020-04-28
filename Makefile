VERSION=$(shell go run main.go -version)

tidy:
	go mod tidy -v
build: tidy
	rm -rf dist
	go run main.go
	./readme.sh $(VERSION)
	git diff HEAD --quiet || exit 1
	tar czvf dist/caddy_$(VERSION)_darwin_amd64.tar.gz dist/caddy_$(VERSION)_darwin_amd64/*
	tar czvf dist/caddy_$(VERSION)_linux_amd64.tar.gz dist/caddy_$(VERSION)_linux_amd64/*
	tar czvf dist/caddy_$(VERSION)_windows_amd64.tar.gz dist/caddy_$(VERSION)_windows_amd64/*
release: build
	hub release create -a dist/caddy_$(VERSION)_darwin_amd64.tar.gz -a dist/caddy_$(VERSION)_linux_amd64.tar.gz -a dist/caddy_$(VERSION)_windows_amd64.tar.gz -m v$(VERSION) v$(VERSION)
