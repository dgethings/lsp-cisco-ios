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

// The Split method needs to provide SplitFunc definition
// type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
func Split(data []byte, _ bool) (adv int, token []byte, err error) {
	header, content, enough := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	// not an error, just don't have enough data yet
	if !enough {
		return 0, nil, nil
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	// we expect to be able to get the content length now
	if err != nil {
		return 0, nil, err
	}

	// not an error, we have valid data just not all of it yet
	if len(content) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}
