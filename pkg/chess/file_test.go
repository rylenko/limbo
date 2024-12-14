package chess

import "testing"

func TestFileString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		file File
		str  string
	}{
		{FileNil, "FileNil"},
		{FileA, "FileA"},
		{FileB, "FileB"},
		{FileC, "FileC"},
		{FileD, "FileD"},
		{FileE, "FileE"},
		{FileF, "FileF"},
		{FileG, "FileG"},
		{FileH, "FileH"},
		{File(123), "<unknown File=123>"},
	}

	for _, test := range tests {
		t.Run(test.str, func(t *testing.T) {
			t.Parallel()

			str := test.file.String()
			if str != test.str {
				t.Fatalf("%s.String() expected %q but got %q", test.str, test.str, str)
			}
		})
	}
}
