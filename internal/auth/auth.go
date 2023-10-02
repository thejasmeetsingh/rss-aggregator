package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIkey Fetch API key from HTTP headers
// Auth format: Authorization: Bearer {API KEY}
func GetAPIkey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("auth credentials is not provided")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 || vals[0] != "Bearer" {
		return "", errors.New("invalid auth format")
	}

	return vals[1], nil
}
