format:
	goimports  -l -w .
build: format
	goreleaser release --snapshot --rm-dist
	./dist/darwin_amd64/caddy -plugins > plugins.txt
	./dist/darwin_amd64/caddy -version > version.txt
publish: build
	goreleaser release --rm-dist
