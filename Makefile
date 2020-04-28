tidy:
	go mod tidy -v
build:
	rm -rf dist
	go run main.go
	./readme.sh
publish: build
