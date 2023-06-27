.DEFAULT_GOAL := build

build:
	go build -o bin/bs-benchmark -ldflags "\
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildVersion=$$(git rev-parse --abbrev-ref HEAD)' \
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildUser=$$(id -u -n)' \
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildTime=$$(date)' \
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildGOOS=$$(go env GOOS)' \
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildARCH=$$(go env GOARCH)' \
		-s -w"

test:
	go test -v ./...
