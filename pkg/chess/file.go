package chess

import "fmt"

// File is the enumeration of all chess board files.
type File uint8

const (
	FileNil File = iota
	FileA
	FileB
	FileC
	FileD
	FileE
	FileF
	FileG
	FileH
)

// Array of all valid files.
var files = [8]File{FileA, FileB, FileC, FileD, FileE, FileF, FileG, FileH}

// String returns string representation of current file.
func (file File) String() string {
	switch file {
	case FileNil:
		return "FileNil"
	case FileA:
		return "FileA"
	case FileB:
		return "FileB"
	case FileC:
		return "FileC"
	case FileD:
		return "FileD"
	case FileE:
		return "FileE"
	case FileF:
		return "FileF"
	case FileG:
		return "FileG"
	case FileH:
		return "FileH"
	default:
		return fmt.Sprintf("<unknown File=%d>", file)
	}
}
