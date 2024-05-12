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
		{input: "enable password", linNum: 0, charNum: 0, expected: ios.EnablePassword},
	}

	for _, tc := range tests {
		actual := selectedWord(tc.input, tc.linNum, tc.charNum)
		assert.Equal(t, tc.expected, actual)
	}
}
