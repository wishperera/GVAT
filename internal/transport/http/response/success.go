package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SuccessResponse struct {
	VATID string `json:"vatId"`
	Valid bool   `json:"valid"`
}

func WriteSuccessResponse(w http.ResponseWriter, res interface{}, headers map[string]string, statusCode int) error {
	for k, v := range headers {
		w.Header().Add(k, v)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if res != nil {
		respByt, err := json.Marshal(res)
		if err != nil {
			return fmt.Errorf("failed to marshal error response due: %s", err)
		}

		_, err = w.Write(respByt)
		if err != nil {
			return fmt.Errorf("failed to write error response due: %s", err)
		}
	}

	return nil
}
