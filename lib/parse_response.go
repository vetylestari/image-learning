package lib

import (
	"io"
	"strings"
)

func ParseResponseToJSON(responseBody io.ReadCloser) (parsed_resp string, err error) {
	buf := new(strings.Builder)
	_, err = io.Copy(buf, responseBody)
	if err != nil {
		return "", err
	}
	parsed_resp = buf.String()
	return parsed_resp, nil
}
