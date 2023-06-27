package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Nathan13888/bs-benchmark/v2/config"
	"github.com/urfave/cli/v2"
)

func startCLI() {
	app := &cli.App{
		Name:    "bs-benchmark",
		Version: config.BuildVersion,
		Usage:   "do you lift bro? bench yo snakes",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "./config.json",
				Usage:       "Load configuration from `FILE`",
				Destination: &config.ConfigFile,
			},
			// TODO: verbose/debug flag
		},
		UsageText: "bs-benchmark [global options] command [command options] [arguments...]",
		ArgsUsage: "[snake1_name snake1_url]...[snakeN_name snakeN_url]",
		Action: func(ctx *cli.Context) error {
			// startup checks
			startupChecks()

			// read flags for snakes
			var snakes []SnakeProp

			// read arguments for snake names and addresses
			for i := 0; i < ctx.Args().Len(); i += 2 {
				name := ctx.Args().Get(i)
				addr := ctx.Args().Get(i + 1)

				if strings.HasPrefix(name, "http") {
					return errors.New("snake name shouldn't start with http (check --help)")
				}
				if !strings.HasPrefix(addr, "http://") && !strings.HasPrefix(addr, "https://") {
					return errors.New("snake address should start with http:// or https:// (check --help)")
				}

				snakes = append(snakes, SnakeProp{
					Name: name,
					Addr: addr,
				})
			}

			if len(snakes) == 0 {
				return errors.New("no snakes provided (check --help)")
			}

			// TODO: ping snakes

			// TODO: auto detect snakes?, --local-detect

			fmt.Println("Running benchmarks...", config.Settings.Rounds, "rounds",
				"and", len(config.Settings.Sizes), "board sizes",
				"with", len(snakes), "snakes.")
			runBenchmarks(&snakes)

			return nil
		},
	}

	app.Suggest = true

	cli.VersionPrinter = func(cCtx *cli.Context) { config.PrintVersion(cCtx) }

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func main() {
	startCLI()
}

func startupChecks() {
	// load configs
	err := config.LoadSettings()
	if err != nil {
		log.Fatal(err)
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

	err = Mkdir(config.Settings.OutputsDir)
	if err != nil {
		log.Fatal(err)
	}
	err = Mkdir(config.Settings.ResultsDir)
	if err != nil {
		log.Fatal(err)
	}
}

func runBenchmarks(snakes *[]SnakeProp) {
	bg := BenchmarkGroup{
		Name:     "default_group_name",
		Rounds:   config.Settings.Rounds,
		Sizes:    config.Settings.Sizes,
		Seed:     config.Settings.Seed,
		Timeout:  config.Settings.Timeout,
		Gametype: config.Settings.Gametype,
		Map:      config.Settings.Map,
		Snakes:   snakes,
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
	resLog, err := bg.WriteJSON()
	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
	}

	fmt.Println("Written results to ", resLog)
}
