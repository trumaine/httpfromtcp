package headers

import (
	"bytes"
	"fmt"
	"strings"
)

const crlf = "\r\n"

type Headers map[string]string

func NewHeaders() Headers {
	return make(Headers)
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return 0, false, nil
	}

	if idx == 0 {
		return 2, true, nil
	}

	headerLineText := string(data[:idx])
	headerKey, headerValue, err := headerLineFromString(headerLineText)
	if err != nil {
		return 0, false, err
	}
	h.Set(headerKey, headerValue)

	return idx + 2, false, nil
}

func headerLineFromString(str string) (key string, value string, err error) {
	parts := strings.SplitN(str, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid header line: %s", str)
	}

	key = parts[0]
	if strings.ContainsAny(key, " \t") {
		return "", "", fmt.Errorf("invalid header name: %s", key)
	}

	value = strings.TrimSpace(parts[1])
	key = strings.TrimSpace(key)

	return key, value, nil
}

func (h Headers) Set(key, value string) {
	h[key] = value
}
