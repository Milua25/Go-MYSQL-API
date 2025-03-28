package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(req *http.Request, X interface{}) {
	if body, err := io.ReadAll(req.Body); err == nil {
		if err := json.Unmarshal(body, X); err != nil {
			return
		}
	}
}
