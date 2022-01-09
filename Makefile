test: 
	go test ./... -v -cover

build:
	GOOS=linux GOARCH=amd64 go build -o bin/sd
	GOOS=darwin GOARCH=amd64 go build -o ./bin/sd_mac
	GOOS=windows GOARCH=amd64 go build -o ./bin/sd_win.exe
	GOOS=linux GOARCH=arm go build -o ./bin/sd_arm
	GOOS=linux GOARCH=arm64 go build -o ./bin/sd_arm64
	GOOS=darwin GOARCH=arm64 go build -o ./bin/sd_arm64_mac

run:
	go run main.go

clean:
	rm -rf ./bin