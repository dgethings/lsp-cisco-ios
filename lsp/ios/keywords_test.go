package ios

import (
	"log"
	"testing"
)

func TestGetKeywords(t *testing.T) {
	tests := map[string][]string{
		"activation-character": {"activation-character ascii-number", "no activation-character"},
		"alias":                {"alias mode command-alias original-command", "no alias mode  [command-alias]"},
	}
	data := getKeywords()
	for kw, expected := range tests {
		actual, ok := data[kw]
		if !ok {
			log.Fatalf("%s keyword not found", kw)
		}
		for i, a := range actual {
			if expected[i] != a {
				log.Fatalf("expected: %s, actual: %s", expected[i], a)
			}
		}
	}
}
