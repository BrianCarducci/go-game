package state

type GameState int

const (
	StateMenu GameState = iota
	StatePlaying
	StatePaused
)