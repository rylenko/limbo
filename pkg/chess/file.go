package chess

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

var (
	// Array of all valid files.
	files = [8]File{FileA, FileB, FileC, FileD, FileE, FileF, FileG, FileH}

	// Mapping of all File variants to strings.
	fileStrings = map[File]string{
		FileNil: "FileNil",
		FileA: "FileA",
		FileB: "FileB",
		FileC: "FileC",
		FileD: "FileD",
		FileE: "FileE",
		FileF: "FileF",
		FileG: "FileG",
		FileH: "FileH",
	}
)

// String returns string representation of current file.
func (file File) String() string {
	str, ok := fileStrings[file]
	if !ok {
		return "<unknown File>"
	}
	return str
}
