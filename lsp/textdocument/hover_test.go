package textdocument

import (
	"testing"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/stretchr/testify/assert"
)

func TestSelectedWord(t *testing.T) {
	type test struct {
		input    string
		expected ios.Keyword
		linNum   uint
		charNum  uint
	}
	tests := []test{
		{input: "enable password", linNum: 0, charNum: 1, expected: ios.EnablePassword},
		{input: "enable secret", linNum: 0, charNum: 1, expected: ios.EnableSecret},
	}

	for _, tc := range tests {
		actual, err := selectedWord(tc.input, tc.linNum, tc.charNum)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, tc.expected, actual)
	}
}
