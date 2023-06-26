package main

import (
	"errors"
	"log"
	"os/exec"
)

// TODO: implement CLI
func main() {
	// TODO: ping snakes

	// TODO: load snakes from config/flags

	runBenchmarks()
}

func runBenchmarks() {
	// resolve BATTLESNAKE_BIN
	path, err := exec.LookPath(BATTLESNAKE_BIN)
	if errors.Is(err, exec.ErrDot) {
		err = nil
	} else if err != nil {
		log.Fatal(err)
	} else {
		BATTLESNAKE_BIN = path
		log.Printf("found BATTLESNAKE_BIN at %s", path)
	}

	for _, size := range SIZES {
		width := size
		height := size

		for round := 0; round < ROUNDS; round++ {
			// create benchmark
			bench := CreateBenchmark(&SNAKES, round, width, height)

			// run benchmark
			bench.Run()
		}
	}
}
