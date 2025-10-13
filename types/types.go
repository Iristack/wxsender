package types

type MessageType int

const (
	APP MessageType = iota
	GroupApp
	ROBOT
)
