hello:
	echo "Hello from Meeseeks"

build:
	go build -o bin/main main.go

run:
	go run main.go


compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o releases/meeseeks-linux-arm 
	GOOS=linux GOARCH=arm64 go build -o releases/meeseeks-linux-arm64 
	GOOS=freebsd GOARCH=386 go build -o releases/meeseeks-freebsd-386 
	GOOS=windows GOARCH=amd64 go build -o releases/meeseeks.exe 

all: hello