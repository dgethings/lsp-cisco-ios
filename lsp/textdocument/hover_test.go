package textdocument

import (
	"fmt"
	"testing"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/stretchr/testify/assert"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
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

var params = new(protocol.HoverParams)

func TestHover(t *testing.T) {
	type test struct {
		params   *protocol.HoverParams
		input    string
		expected string
	}
	tests := []test{
		{
			input:    "enable password 15 7 $1$1abcdef0123456789",
			expected: "To set a local password to control access to various privilege levels",
			params:   params,
		},
		{
			input:    "enable secret 15 0 abcdef0123456789",
			expected: "To specify an additional layer of security over the enable password command",
			params:   params,
		},
	}
	for _, tc := range tests {
		State[""] = tc.input
		actual, err := Hover(&glsp.Context{}, tc.params)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println("Actual: ", actual.Contents)
		fmt.Println("Expected: ", tc.expected)
		assert.Contains(t, actual.Contents, tc.expected)
	}
}
