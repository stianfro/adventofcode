generate:
	buf generate

run:
	air --build.cmd "go build -o bin/server server/main.go" --build.bin "./bin/server"

ui:
	grpcui --port 56674 --plaintext localhost:8080

