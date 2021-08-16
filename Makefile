test: 
	go test ./...

build:
	go build -o bin/app

clean:
	rm -rf ./bin