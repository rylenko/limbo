package chess

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

var files = [8]File{FileA, FileB, FileC, FileD, FileE, FileF, FileG, FileH}
