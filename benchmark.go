package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

type Benchmark struct {
	Args    []string
	Command string `json:"command"`
	LogFile string `json:"log_file"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Snakes  *[]Snake
}

type BenchmarkResult struct {
	Bench *Benchmark
	// TODO: who wins? parse log file
}

type Snake struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
}

func CreateBenchmark(snakes *[]Snake, round int, width int, height int) Benchmark {
	var (
		gametype = GAMETYPE
		mapp     = MAP
		timeout  = TIMEOUT
		seed     = SEED
	)

	size := fmt.Sprintf("%dx%d", width, height)
	logFile := fmt.Sprintf("benchmark-%s-r%d-%s-%s-%s.json",
		size, round, gametype, mapp, seed)

	frmt := []string{
		"play",
		"--width", fmt.Sprintf("%d", width), "--height", fmt.Sprintf("%d", height),
		"--timeout", timeout,
		"--gametype", gametype,
		"--map", mapp,
		"--seed", seed,
		"--output", logFile,
	}

	if USE_BROWSER {
		frmt = append(frmt, "--browser")
	}

	for _, snake := range *snakes {
		frmt = append(frmt, "--name", snake.Name, "--url", snake.Addr)
	}

	res := Benchmark{
		Args:    frmt,
		Command: fmt.Sprintf("%s %s", BATTLESNAKE_BIN, strings.Join(frmt, " ")),
		LogFile: logFile,
		Width:   width,
		Height:  height,
		Snakes:  snakes,
	}
	return res
}

func (benchmark Benchmark) Run() BenchmarkResult {
	args := benchmark.Args

	log.Println("Running BATTLESNAKE_BIN with args:", args)

	cmd := exec.Command(BATTLESNAKE_BIN, args...)

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
		Bench: &benchmark,
	}
}
