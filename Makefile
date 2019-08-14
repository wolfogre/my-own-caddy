tidy:
	rm -rf go.mod go.sum
	go mod init caddy
	go mod tidy -v
build:
	goreleaser release --parallelism=12 --snapshot --rm-dist
	./dist/caddy_darwin_amd64/caddy -version > README.txt
	echo "" >> README.txt
	./dist/caddy_darwin_amd64/caddy -plugins >> README.txt
publish: build
	goreleaser release --parallelism=12 --rm-dist
