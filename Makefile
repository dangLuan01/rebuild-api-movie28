server: 
	cd ./cmd/api && go run .
host:
	cd ./cmd/api && ./api
build:
	cd ./cmd/api && go build -mod=vendor -o api .
		