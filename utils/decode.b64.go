package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"strings"
)

func GetB64Image(base64str string) (io.Reader, error) {
	parts := strings.SplitN(base64str, ",", 2)
	if len(parts) != 2 {
		return nil, errors.New("invalid string")
	}
	data, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
