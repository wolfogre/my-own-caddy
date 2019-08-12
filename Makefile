format:
	goimports -l -w .
build: format
	goreleaser release --parallelism=12 --snapshot --rm-dist
	./dist/caddy_darwin_amd64/caddy -version > README.txt
	echo "" >> README.txt
	./dist/caddy_darwin_amd64/caddy -plugins >> README.txt
publish: build
	goreleaser release --parallelism=12 --rm-dist
