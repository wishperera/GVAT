package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Exception struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	TraceID     string `json:"trace_id"`
}

func WriteExceptionResponse(w http.ResponseWriter, e Exception, headers map[string]string, statusCode int) error {
	for k, v := range headers {
		w.Header().Add(k, v)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode == http.StatusMethodNotAllowed {
		return nil
	}

	respByt, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("failed to marshal error response due: %s", err)
	}

	_, err = w.Write(respByt)
	if err != nil {
		return fmt.Errorf("failed to write error response due: %s", err)
	}

	return nil
}
