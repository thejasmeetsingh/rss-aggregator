package utils

import (
	"net/http"
	"strconv"
)

func GetOffset(r *http.Request) int32 {
	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	return int32(offset)
}
