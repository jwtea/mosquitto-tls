package main

const (
	StateAction string = "STATE"
	ModeAction  string = "MODE"
)

type ModeRequest struct {
	Timestamp int64
	Mode      *Mode
}

type StateRequest struct {
	Timestamp int64
	State     State
}

// Mode types
type Mode int32

const (
	ModeAuto   Mode = 0
	ModeManual Mode = 1
)

// State Types
type State int32

const (
	StateOFF State = 0
	StateON  State = 1
)

type Topic string

const (
	StateTopic Topic = "state"
	ModeTopic  Topic = "mode"
)
