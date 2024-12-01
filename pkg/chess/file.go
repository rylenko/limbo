package chess

const filesCount = 8

// File represents chess files: from A to H.
//
// FileA has the lowest value, FileH has the highest value.
type File uint8

const (
	FileA File = iota
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)
