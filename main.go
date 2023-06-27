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

	// load configs
	err := config.LoadSettings()
	if err != nil {
		log.Fatal(err)
	}

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
	config.Settings.OutputsDir = path
	// fmt.Println("OUTPUTS_DIR:", config.OUTPUTS_DIR)

	// prepare ./results directory
	res_path, err := filepath.Abs("./results")
	if err != nil {
		log.Fatal(err)
	}
	config.Settings.ResultsDir = res_path

	runBenchmarks()
}

func runBenchmarks() {
	bg := BenchmarkGroup{
		Name:     "default_group_name",
		Rounds:   config.Settings.Rounds,
		Sizes:    config.Settings.Sizes,
		Seed:     config.Settings.Seed,
		Timeout:  config.Settings.Timeout,
		Gametype: config.Settings.Gametype,
		Map:      config.Settings.Map,
		Snakes: &[]SnakeProp{ // TODO: load from config
			{"rng0", "http://127.0.0.1:8000"},
			{"rng1", "http://127.0.0.1:8001"},
			{"rng2", "http://127.0.0.1:8002"},
			{"rng3", "http://127.0.0.1:8003"},
		},
	}

	// resolve BATTLESNAKE_BIN
	path, err := exec.LookPath(config.Settings.BATTLESNAKE_BIN)
	if errors.Is(err, exec.ErrDot) {
		err = nil
	} else if err != nil {
		log.Fatal(err)
	} else {
		config.Settings.BATTLESNAKE_BIN = path
		// log.Printf("found BATTLESNAKE_BIN at %s", path)
	}

	// results := make([]BenchmarkResult, len(bg.Sizes)*bg.Rounds)
	var results []BenchmarkResult
	bg.Benchmarks = &results

	draws := 0
	wins := make(map[string]int)

	for _, size := range bg.Sizes {
		width := size
		height := size

		for round := 0; round < bg.Rounds; round++ {
			// create benchmark
			bench := bg.CreateBenchmark(round, width, height)

			// run benchmark
			res := bench.Run()
			// results[i*len(bg.Sizes)+round] = res
			results = append(results, res)

			data := res.Bench.ParseLog()
			if data != nil {
				draw := data.Result.IsDraw
				winner := data.Result.WinnerName

				if draw {
					draws++
				} else {
					wins[winner]++
				}
			}
		}
	}

	summary := BenchmarkSummary{
		Draws: draws,
		Wins:  wins,
	}
	bg.Summary = &summary

	// write JSON to file to ./results
	err = bg.WriteJSON()
}
