VERSION=$(shell go run main.go -version)

tidy:
	go mod tidy -v
build: tidy
	rm -rf dist
	go run main.go
	./readme.sh $(VERSION)
	git diff HEAD --quiet || exit 1
	file dist/caddy_$(VERSION)_darwin_amd64/caddy | grep -q "Mach-O 64-bit executable"
	file dist/caddy_$(VERSION)_linux_amd64/caddy | grep "ELF 64-bit LSB executable"
	file dist/caddy_$(VERSION)_windows_amd64/caddy.exe | grep "PE32+ executable"
	cd dist/caddy_$(VERSION)_darwin_amd64 && tar czvf ../caddy_$(VERSION)_darwin_amd64.tar.gz *
	cd dist/caddy_$(VERSION)_linux_amd64 && tar czvf ../caddy_$(VERSION)_linux_amd64.tar.gz *
	cd dist/caddy_$(VERSION)_windows_amd64 && tar czvf ../caddy_$(VERSION)_windows_amd64.tar.gz *
release: build
	hub release create -a dist/caddy_$(VERSION)_darwin_amd64.tar.gz -a dist/caddy_$(VERSION)_linux_amd64.tar.gz -a dist/caddy_$(VERSION)_windows_amd64.tar.gz -m v$(VERSION) v$(VERSION)
