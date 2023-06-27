# Battlesnake Benchmark

## Usage

**Test everything with Docker**: `docker compose up -d && docker compose logs -f`

**Start dummy snakes**: `docker compose start rng0 rng1 rng2 rng3`

**Start local board**: `docker compose start board`

**Build Binary (local)**: `make` (builds to `./bin`)

**CLI Help**: `bs-benchmark`


## Config

Logs output to `./outputs/` by default.

