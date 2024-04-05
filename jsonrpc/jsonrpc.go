package jsonrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type BaseMessage struct {
	Method string `json:"method"`
}

func Encode(msg any) string {
	contents, err := json.Marshal(msg)
	// TODO: handle this error better
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(contents), contents)
}

func Decode(msg []byte) (string, []byte, error) {
	header, content, ok := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	var baseMessage BaseMessage
	if !ok {
		return "", nil, errors.New("did not find separator")
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}
