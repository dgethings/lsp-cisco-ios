package jsonrpc_test

import (
	"testing"

	"github.com/dgethings/lsp-cisco-ios/jsonrpc"
)

type EncodeExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := jsonrpc.Encode(EncodeExample{Testing: true})
	if actual != expected {
		t.Fatalf("expected %s, actual %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	expected := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := jsonrpc.Decode([]byte(expected))
	length := len(content)
	if err != nil {
		t.Fatal(err)
	}
	if length != 15 {
		t.Fatalf("expected 15, actual %d", length)
	}
	if method != "hi" {
		t.Fatalf("expected 'hi', actual '%s'", method)
	}
}
