package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/Nathan13888/bs-benchmark/v2/config"
)

type Benchmark struct {
	// TODO: UUID
	Args    []string        `json:"args"`
	Command string          `json:"command"`
	LogFile string          `json:"log_file"`
	Width   int             `json:"width"`
	Height  int             `json:"height"`
	Snakes  *[]SnakeProp    `json:"-"`
	Group   *BenchmarkGroup `json:"-"`
}

type BenchmarkGroup struct {
	Name       string             `json:"name"`
	Rounds     int                `json:"rounds"`
	Sizes      []int              `json:"sizes"`
	Seed       string             `json:"seed"`
	Timeout    int                `json:"timeout"`
	Gametype   string             `json:"gametype"`
	Map        string             `json:"map"`
	Snakes     *[]SnakeProp       `json:"snakes"`
	Benchmarks *[]BenchmarkResult `json:"benchmarks"`
	Summary    *BenchmarkSummary  `json:"summary"`
}

type BenchmarkSummary struct {
	Draws int            `json:"draws"`
	Wins  map[string]int `json:"wins"`
}

type BenchmarkResult struct {
	Bench *Benchmark `json:"benchmark"`
	// TODO: who wins? parse log file
}

type SnakeProp struct {
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
	currentTime := GetShortTime()
	logFile := fmt.Sprintf("benchmark-%s-%s-r%d-%s-%s-%s.json",
		currentTime, size, round, gametype, mapp, seed)

	frmt := []string{
		"play",
		"--width", fmt.Sprintf("%d", width), "--height", fmt.Sprintf("%d", height),
		"--timeout", fmt.Sprintf("%d", timeout),
		"--gametype", gametype,
		"--map", mapp,
		"--seed", seed,
		"--output", filepath.Join(config.Settings.OutputsDir, logFile),
		"--board-url", config.Settings.BoardURL,
	}

	if config.Settings.UseBrowser {
		frmt = append(frmt, "--browser")
	}

	for _, snake := range *bg.Snakes {
		frmt = append(frmt, "--name", snake.Name, "--url", snake.Addr)
	}

	res := Benchmark{
		Args:    frmt,
		Command: fmt.Sprintf("%s %s", config.Settings.BATTLESNAKE_BIN, strings.Join(frmt, " ")),
		LogFile: logFile,
		Width:   width,
		Height:  height,
		Snakes:  bg.Snakes,
	}
	return res
}

func (benchmark *Benchmark) Run() BenchmarkResult {
	args := benchmark.Args

	cmd := exec.Command(config.Settings.BATTLESNAKE_BIN, args...)

	// stdout, err := cmd.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(stdout)

	return BenchmarkResult{
		Bench: benchmark,
	}
}

type LogData struct {
	Game     Game
	Requests []SnakeRequest
	Result   Result
}

func (b *Benchmark) ParseLog() *LogData {
	logFile := b.LogFile
	path := filepath.Join(config.Settings.OutputsDir, logFile)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data LogData
	raw := strings.Split(string(content), "\n")

	var lines []string
	for _, line := range raw {
		if line != "" {
			lines = append(lines, line)
		}
	}

	n := len(lines)
	if n <= 3 {
		log.Fatal("Log file is too short.")
	}

	err = json.Unmarshal([]byte(lines[0]), &data.Game)
	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i < n-1; i++ {
		var req SnakeRequest
		err = json.Unmarshal([]byte(lines[i]), &req)
		if err != nil {
			fmt.Println(err)
		}
		data.Requests = append(data.Requests, req)
	}

	err = json.Unmarshal([]byte(lines[n-1]), &data.Result)
	if err != nil {
		fmt.Println(err)
	}

	return &data
}

func (bg *BenchmarkGroup) EncodeJSON() []byte {
	// encoded, err := json.MarshalIndent(bg, "", "  ")
	encoded, err := json.MarshalIndent(bg, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return encoded
}

// write JSON to file in ./results/
func (bg *BenchmarkGroup) WriteJSON() error {
	resultName := fmt.Sprintf("results-%s-%s-%s-%s.json", GetShortTime(), bg.Name, bg.Gametype, bg.Map)
	file := filepath.Join(config.Settings.ResultsDir, resultName)
	resultFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	formattedOutput := string(bg.EncodeJSON())

	for _, line := range formattedOutput {
		_, err := io.WriteString(resultFile, fmt.Sprintf("%s\n", line))
		if err != nil {
			return err
		}
	}

	return nil
}

// Print all contents of BenchmarkGroup by encoding JSON
func (bg *BenchmarkGroup) PrintJSON() {
	fmt.Println(string(bg.EncodeJSON()))
}
