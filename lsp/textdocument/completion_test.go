package textdocument

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestCompletion(t *testing.T) {
	type test struct {
		params   *protocol.CompletionParams
		input    string
		expected []string
	}
	tests := []test{
		{
			input: "en",
			expected: []string{"enable"},
			params: &protocol.CompletionParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					Position: protocol.Position{
						Line:      0,
						Character: 2,
					},
				},
			},
		},
		{
			input: "enable ",
			expected: []string{"password", "secret"},
			params: &protocol.CompletionParams{
				TextDocumentPositionParams: protocol.TextDocumentPositionParams{
					Position: protocol.Position{
						Line:      0,
						Character: 7,
					},
				},
			},
		},
	}

	for _, tc := range tests {
		State[""] = tc.input
		actual, err := Completion(&glsp.Context{}, tc.params)
		if err != nil {
			t.Fatal(err)
		}

		actualCompletions := actual.([]protocol.CompletionItem)
		var actualLabels []string
		for _, c := range actualCompletions {
			actualLabels = append(actualLabels, c.Label)
		}

		assert.ElementsMatch(t, tc.expected, actualLabels)
	}
}
