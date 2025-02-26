BUILD_DIR=./build

clean:
	rm -rf ${BUILD_DIR}

build:
	if [ '$(shell echo "${GIT_TAG}" | cut -c 1 )' != 'v' ]; then exit 1; fi;
	$(eval LD_FLAGS := -X main.Version=$(shell echo ${GIT_TAG} | cut -c 2-) -X main.BuildDate=$(shell date "+%F-%T") -X main.Commit=$(shell git rev-parse --verify HEAD))
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64       go build -ldflags="${LD_FLAGS}" -o build/wb2-cli-windows-amd64.exe wb2-cli.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386         go build -ldflags="${LD_FLAGS}" -o build/wb2-cli-windows-386.exe   wb2-cli.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64       go build -ldflags="${LD_FLAGS}" -o build/wb2-cli-linux-amd64       wb2-cli.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=386         go build -ldflags="${LD_FLAGS}" -o build/wb2-cli-linux-386         wb2-cli.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64       go build -ldflags="${LD_FLAGS}" -o build/wb2-cli-linux-arm64       wb2-cli.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm GOARM=7 go build -ldflags="${LD_FLAGS}" -o build/wb2-cli-linux-arm-7       wb2-cli.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64       go build -ldflags="${LD_FLAGS}" -o build/wb2-cli-darwin-amd64      wb2-cli.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64       go build -ldflags="${LD_FLAGS}" -o build/wb2-cli-darwin-arm64      wb2-cli.go
