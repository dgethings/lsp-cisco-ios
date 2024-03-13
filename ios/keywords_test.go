package ios

import (
	"log"
	"testing"
)

func TestGetKeywords(t *testing.T) {
	data := getKeywords()
	log.Println(data)
	t.Fatal(data)
}
