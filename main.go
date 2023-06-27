package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Nathan13888/bs-benchmark/v2/config"
)

// TODO: implement CLI
func main() {
	// TODO: ping snakes

	// TODO: load snakes from config/flags

	// prepare ./outputs directory
	if _, err := os.Stat("./outputs"); os.IsNotExist(err) {
		// create directory
		err := os.Mkdir("./outputs", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	// get absolute path of ./outputs directory
	path, err := filepath.Abs("./outputs")
	if err != nil {
		log.Fatal(err)
	}
	config.OUTPUTS_DIR = path
	// fmt.Println("OUTPUTS_DIR:", config.OUTPUTS_DIR)

	runBenchmarks()
}

func runBenchmarks() {
	bg := BenchmarkGroup{
		Name:     "default_group_name",
		Rounds:   config.ROUNDS,
		Sizes:    config.SIZES,
		Seed:     config.SEED,
		Timeout:  config.TIMEOUT,
		Gametype: config.GAMETYPE,
		Map:      config.MAP,
		Snakes: &[]Snake{ // TODO: load from config
			{"rng0", "http://127.0.0.1:8000"},
			{"rng1", "http://127.0.0.1:8001"},
			{"rng2", "http://127.0.0.1:8002"},
			{"rng3", "http://127.0.0.1:8003"},
		},
	}

	// resolve BATTLESNAKE_BIN
	path, err := exec.LookPath(config.BATTLESNAKE_BIN)
	if errors.Is(err, exec.ErrDot) {
		err = nil
	} else if err != nil {
		log.Fatal(err)
	} else {
		config.BATTLESNAKE_BIN = path
		// log.Printf("found BATTLESNAKE_BIN at %s", path)
	}

	results := make([]BenchmarkResult, len(bg.Sizes)*bg.Rounds)
	bg.Benchmarks = &results

	for i, size := range bg.Sizes {
		width := size
		height := size

		for round := 0; round < bg.Rounds; round++ {
			// create benchmark
			bench := bg.CreateBenchmark(round, width, height)

			// run benchmark
			res := bench.Run()
			results[i*len(bg.Sizes)+round] = res
		}
	}

	// bg.PrintJSON()
}
