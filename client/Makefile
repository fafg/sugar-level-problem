build:
	go build -ldflags "-s -w" -o client-test main.go
	ls -lah client-test

test:
	@ docker-compose up -d && echo "sleeping 3s to make sure container is up.\n" && sleep 3s
	@ go test -v -cover ./...
	@ docker-compose kill