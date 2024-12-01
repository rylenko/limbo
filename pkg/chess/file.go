package chess

const filesCount = 8

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
