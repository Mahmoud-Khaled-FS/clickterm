package utils

import (
	"encoding/base64"
	"errors"
)

func EncodeStr(str string) string {
	encoding := base64.StdEncoding.EncodeToString([]byte(str))
	return encoding
}

func DecodeStr(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", errors.New("can not decode string")
	}
	return string(data), nil
}
