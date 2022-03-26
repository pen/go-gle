package util

import (
	"bytes"
	"compress/bzip2"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/ulikunitz/xz"
)

func GetEncodedEnv(name string) ([]byte, error) {
	mr, err := NewMultiEncodedStringReader(os.Getenv(name))
	if err != nil {
		return nil, fmt.Errorf("on decoder: %w", err)
	}

	b, err := io.ReadAll(mr)
	if err != nil {
		return nil, fmt.Errorf("on ReadAll(): %w", err)
	}

	return b, nil
}

func NewMultiEncodedStringReader(str string) (io.Reader, error) {
	if len(str) < 2 {
		return nil, fmt.Errorf("on NewMultiEncodedStringReader(): too short")
	}

	var reader io.Reader = bytes.NewReader([]byte(str))
	if str[0] == '{' {
		return reader, nil
	}

	reader = base64.NewDecoder(base64.StdEncoding, reader)

	switch str[0:2] {
	case "Ql":
		reader = bzip2.NewReader(reader)
	case "/T":
		var err error

		reader, err = xz.NewReader(reader)
		if err != nil {
			return nil, fmt.Errorf("on xz.NewReader(): %w", err)
		}
	}

	return reader, nil
}

func StripQuotes(str string) string {
	if l := len(str); l >= 2 {
		if q := str[0]; q == '"' || q == '\'' {
			if q == str[l-1] {
				return str[1 : l-1]
			}
		}
	}

	return str
}
