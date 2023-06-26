package main

// settings
var (
	BATTLESNAKE_BIN = "battlesnake"
	SNAKES          = []Snake{
		{"rng0", "http://127.0.0.1:8000"},
		{"rng1", "http://127.0.0.1:8001"},
		{"rng2", "http://127.0.0.1:8002"},
		{"rng3", "http://127.0.0.1:8003"},
	}
	USE_BROWSER = true

	ROUNDS int = 50
	SIZES      = []int{5, 7, 9, 11}

	SEED     string = "1656460409268690000"
	TIMEOUT         = "500"
	GAMETYPE        = "standard"
	MAP             = "standard"
)

// TODO: config loading function
