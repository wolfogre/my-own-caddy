format:
	goimports  -l -w .
build: format
	goreleaser release --snapshot --rm-dist
publish: format
	goreleaser release --rm-dist
