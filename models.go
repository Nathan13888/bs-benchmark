package main

import "time"

// from outputs.go
type GameExporter struct {
	game          Game
	snakeRequests []SnakeRequest
	winner        SnakeState
	isDraw        bool
}

type Result struct {
	WinnerID   string `json:"winnerId"`
	WinnerName string `json:"winnerName"`
	IsDraw     bool   `json:"isDraw"`
}

// from play.go

// Used to store state for each SnakeState while running a local game
type SnakeState struct {
	URL        string
	Name       string
	ID         string
	LastMove   string
	Character  rune
	Color      string
	Head       string
	Tail       string
	Author     string
	Version    string
	Error      error
	StatusCode int
	Latency    time.Duration
}

// from models.go
type SnakeRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

// from models.go

type Ruleset struct {
	Name     string          `json:"name"`
	Version  string          `json:"version"`
	Settings RulesetSettings `json:"settings"`
}

// Game represents the current game state
type Game struct {
	ID      string  `json:"id"`
	Ruleset Ruleset `json:"ruleset"`
	Map     string  `json:"map"`
	Timeout int     `json:"timeout"`
	Source  string  `json:"source"`
}

// Board provides information about the game board
type Board struct {
	Height  int     `json:"height"`
	Width   int     `json:"width"`
	Snakes  []Snake `json:"snakes"`
	Food    []Coord `json:"food"`
	Hazards []Coord `json:"hazards"`
}

// Snake represents information about a snake in the game
type Snake struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	Latency        string         `json:"latency"`
	Health         int            `json:"health"`
	Body           []Coord        `json:"body"`
	Head           Coord          `json:"head"`
	Length         int            `json:"length"`
	Shout          string         `json:"shout"`
	Squad          string         `json:"squad"`
	Customizations Customizations `json:"customizations"`
}

type Customizations struct {
	Color string `json:"color"`
	Head  string `json:"head"`
	Tail  string `json:"tail"`
}

// Coord represents a point on the board
type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// RulesetSettings contains a static collection of a few settings that are exposed through the API.
type RulesetSettings struct {
	FoodSpawnChance     int            `json:"foodSpawnChance"`
	MinimumFood         int            `json:"minimumFood"`
	HazardDamagePerTurn int            `json:"hazardDamagePerTurn"`
	HazardMap           string         `json:"hazardMap"`       // Deprecated, replaced by Game.Map
	HazardMapAuthor     string         `json:"hazardMapAuthor"` // Deprecated, no planned replacement
	RoyaleSettings      RoyaleSettings `json:"royale"`
	SquadSettings       SquadSettings  `json:"squad"` // Deprecated, provided with default fields for API compatibility
}

// RoyaleSettings contains settings that are specific to the "royale" game mode
type RoyaleSettings struct {
	ShrinkEveryNTurns int `json:"shrinkEveryNTurns"`
}

// SquadSettings contains settings that are specific to the "squad" game mode
type SquadSettings struct {
	AllowBodyCollisions bool `json:"allowBodyCollisions"`
	SharedElimination   bool `json:"sharedElimination"`
	SharedHealth        bool `json:"sharedHealth"`
	SharedLength        bool `json:"sharedLength"`
}
