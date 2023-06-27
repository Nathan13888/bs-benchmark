.DEFAULT_GOAL := build

build:
	go build -o bin/bs-benchmark -ldflags "\
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildVersion=$$(git rev-parse --abbrev-ref HEAD)' \
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildUser=$$(id -u -n)' \
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildTime=$$(date)' \
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildGOOS=$$(go env GOOS)' \
		-X 'github.com/Nathan13888/bs-benchmark/v2/config.BuildARCH=$$(go env GOARCH)' \
		-s -w"

dcu:
	docker compose up -d
	docker compose logs -f

dc-snakes:
	docker compose up rng0 rng1 rng2 rng3

dc-board:
	docker compose up board

demo:
	docker compose up -d rng0 rng1 rng2 rng3
	cat config.json
	go run . "rng0" "http://127.0.0.1:8000" "rng1" "http://127.0.0.1:8001" "rng2", "http://127.0.0.1:8002" "rng3" "http://127.0.0.1:8003"

test:
	go test -v ./...
