package textdocument

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

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
			params: &protocol.HoverParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					Position: protocol.Position{
						Line:      0,
						Character: 8,
					},
				},
			},
		},
		{
			input:    "enable secret 15 0 abcdef0123456789",
			expected: "To specify an additional layer of security over the enable password command",
			params: &protocol.HoverParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					Position: protocol.Position{

						Line:      0,
						Character: 8,
					},
				},
			},
		},
		{
			input:    "enable password 15 user 9 abc123\nenable secret 15 9 abc123",
			expected: "To specify an additional layer of security over the enable password command",
			params: &protocol.HoverParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					Position: protocol.Position{
						Line:      1,
						Character: 8,
					},
				},
			},
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