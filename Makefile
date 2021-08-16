test: 
	go test ./... -v -cover

build:
	go build -o bin/app

clean:
	rm -rf ./bin