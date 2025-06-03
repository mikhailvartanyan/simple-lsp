package rpc_test

import "testing"
import "simple-lsp/rpc"

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"yo\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	if err != nil {
		t.Fatal(err)
	}

	contentLength := len(content)
	if contentLength != 15 {
		t.Fatalf("Expected: %d, Actual: %d", 16, contentLength)
	}
	if method != "yo" {
		t.Fatalf("Expected: %s, Actual: %s", "yo", method)
	}
}
