# Battlesnake Benchmark

## Usage

**Command Usage Information**: `bs-battlesnake --help`

**Test everything with Docker**: `docker compose up -d && docker compose logs -f`

**Start dummy snakes**: `docker compose start rng0 rng1 rng2 rng3`

**Start local board**: `docker compose start board`

**Build Binary (local)**: `make` (builds to `./bin`)

**Demo**: `make demo`

## Config

Config is located at `./config.json` by default but could be configured with `--config` (or `-c`)

Logs output to `./outputs/` by default.

## Todo. Progress Map
- [ ] Elo/TrueSkill system
- [ ] Game database
- [ ] Remote game collection server
- [ ] More benchmark statistics?
  - [ ] Time to completion
  - [ ] Game stats (eg. length/duration)
- [ ] Parallelized testing/benchmarking? (through workers)

