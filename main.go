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

	// create benchmark
	bench := CreateBenchmark(&SNAKES, 0, SIZES[0], SIZES[0])

	// run benchmark
	bench.Run()

}
