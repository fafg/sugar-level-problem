build:
	go build -ldflags "-s -w" -o client-test main.go

test:
	@ docker-compose up -d && echo "sleeping 3s to make sure container is up.\n" && sleep 3s
	@ pushd client && go test -v -cover && popd && echo "\n"
	@ pushd decisiontree && go test -v -cover && popd && echo "\n"
	@ pushd reports && go test -v -cover && popd && echo "\n"
	@ docker-compose kill