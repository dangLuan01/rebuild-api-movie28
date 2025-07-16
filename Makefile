server: 
	cd ./cmd/api && go run .
host:
	cd ./cmd/api && main.exe
build:
	cd ./cmd/api && go build -o main.exe .
		