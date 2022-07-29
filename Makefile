install:
	go mod tidy
clean:
	rm -rf bin
build:
	go build -o bin/chat
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/chat-freebsd-386
	GOOS=linux GOARCH=386 go build -o bin/chat-linux-386
	GOOS=windows GOARCH=386 go build -o bin/chat-windows-386
run:
	go run .
