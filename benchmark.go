package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Nathan13888/bs-benchmark/v2/config"
)

type Benchmark struct {
	// TODO: UUID
	Args    []string        `json:"args"`
	Command string          `json:"command"`
	LogFile string          `json:"log_file"`
	Width   int             `json:"width"`
	Height  int             `json:"height"`
	Snakes  *[]Snake        `json:"-"`
	Group   *BenchmarkGroup `json:"-"`
}

type BenchmarkGroup struct {
	Name       string             `json:"name"`
	Rounds     int                `json:"rounds"`
	Sizes      []int              `json:"sizes"`
	Seed       string             `json:"seed"`
	Timeout    string             `json:"timeout"`
	Gametype   string             `json:"gametype"`
	Map        string             `json:"map"`
	Snakes     *[]Snake           `json:"snakes"`
	Benchmarks *[]BenchmarkResult `json:"benchmarks"`
}

type BenchmarkResult struct {
	Bench *Benchmark `json:"benchmark"`
	// TODO: who wins? parse log file
}

type Snake struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func (bg *BenchmarkGroup) CreateBenchmark(round int, width int, height int) Benchmark {
	var (
		gametype = bg.Gametype
		mapp     = bg.Map
		timeout  = bg.Timeout
		seed     = bg.Seed
	)

	size := fmt.Sprintf("%dx%d", width, height)
	currentTime := time.Now().Format("20060102150405")
	logFile := fmt.Sprintf("benchmark-%s-%s-r%d-%s-%s-%s.json",
		currentTime, size, round, gametype, mapp, seed)

	frmt := []string{
		"play",
		"--width", fmt.Sprintf("%d", width), "--height", fmt.Sprintf("%d", height),
		"--timeout", timeout,
		"--gametype", gametype,
		"--map", mapp,
		"--seed", seed,
		"--output", filepath.Join(config.OUTPUTS_DIR, logFile),
	}

	if config.USE_BROWSER {
		frmt = append(frmt, "--browser")
	}

	for _, snake := range *bg.Snakes {
		frmt = append(frmt, "--name", snake.Name, "--url", snake.Addr)
	}

	res := Benchmark{
		Args:    frmt,
		Command: fmt.Sprintf("%s %s", config.BATTLESNAKE_BIN, strings.Join(frmt, " ")),
		LogFile: logFile,
		Width:   width,
		Height:  height,
		Snakes:  bg.Snakes,
	}
	return res
}

func (benchmark *Benchmark) Run() BenchmarkResult {
	args := benchmark.Args

	// log.Println("Running BATTLESNAKE_BIN with args:", args)

	cmd := exec.Command(config.BATTLESNAKE_BIN, args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(stdout)

	// TODO: process output results

	return BenchmarkResult{
		Bench: benchmark,
	}
}

// Print all contents of BenchmarkGroup by encoding JSON
func (bg *BenchmarkGroup) PrintJSON() {
	encoded, err := json.MarshalIndent(bg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(encoded))
}
