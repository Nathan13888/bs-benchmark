package config

// settings
var (
	BATTLESNAKE_BIN = "battlesnake"
	OUTPUTS_DIR     = "./outputs"
	USE_BROWSER     = true
	BOARD_URL       = "http://localhost:3000" // https://board.battlesnake.com

	ROUNDS int = 10
	SIZES      = []int{5, 7, 9, 11}

	SEED     string = "1656460409268690000"
	TIMEOUT         = "500"
	GAMETYPE        = "standard"
	MAP             = "standard"
)

// TODO: config loading function
