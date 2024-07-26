package textdocument

import (
	"fmt"
	"testing"

	"github.com/dgethings/lsp-cisco-ios/lsp/ios"
	"github.com/stretchr/testify/assert"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func TestLongestMatch(t *testing.T) {
	tests := []struct {
		line     string
		keywords []ios.Keyword
		expected ios.Keyword
	}{
		{
			"a",
			[]ios.Keyword{ios.NewKeyword("a", "a", "a"), ios.NewKeyword("a b", "a b", "a b"), ios.NewKeyword("a b c", "a b c", "a b c")},
			ios.NewKeyword("a", "a", "a"),
		},
		{
			"a b",
			[]ios.Keyword{ios.NewKeyword("a", "a", "a"), ios.NewKeyword("a b", "a b", "a b"), ios.NewKeyword("a b c", "a b c", "a b c")},
			ios.NewKeyword("a b", "a b", "a b"),
		},
		{
			"a b c",
			[]ios.Keyword{ios.NewKeyword("a", "a", "a"), ios.NewKeyword("a b", "a b", "a b"), ios.NewKeyword("a b c", "a b c", "a b c")},
			ios.NewKeyword("a b c", "a b c", "a b c"),
		},
	}

	for _, tc := range tests {
		actual := longestMatch(tc.line, tc.keywords)
		assert.Equal(t, tc.expected, actual, "longestMatch does not work")
	}
}

func TestNextSpaceIndex(t *testing.T) {
	tests := []struct {
		line     string
		colNum   int
		expected int
	}{
		{"a", 1, 1},
		{"a b", 1, 2},
		{"ab c", 1, 3},
		{"ab c d", 4, 5},
		{"ab c d", 3, 5},
	}

	for _, tc := range tests {
		actual := nextSpaceIndex(tc.line, tc.colNum)
		assert.Equal(t, tc.expected, actual, "line: %s", tc.line)
	}
}

func TestContentAtLine(t *testing.T) {
	tests := []struct {
		contents string
		lineNum  int
		expected string
	}{
		{"a b c\n", 1, "a b c"},
		{"a\nb\nc\n", 1, "a"},
		{"a\nb\nc\n", 2, "b"},
		{"a\nb\nc\n", 3, "c"},
	}

	for _, tc := range tests {
		actual := contentAtLine(tc.contents, tc.lineNum)
		assert.Equal(t, tc.expected, actual, "contents: %s", tc.contents)
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
		{
			input:    "enable password 15 user 9 abc123\nenable secret 15 9 abc123",
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
