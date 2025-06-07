package types

type MainAxisAlignment int
type MainAxisSize int

const (
	MainAxisAlignmentStart MainAxisAlignment = iota
	MainAxisAlignmentCenter
	MainAxisAlignmentEnd
	MainAxisAlignmentSpaceBetween
	MainAxisAlignmentSpaceAround
	MainAxisAlignmentSpaceEvenly
)

const (
	MainAxisSizeMin MainAxisSize = iota // Takes minimum size needed for children
	MainAxisSizeMax                     // Takes maximum size (parent size)
)
