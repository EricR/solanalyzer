all: test build
build: 
	dep ensure
	mkdir -p build
	go build -o build/solanalyzer -v
test: 
	go test -v ./...
clean: 
	go clean
	rm -f build/solanalyzer