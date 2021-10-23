package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteSuccessResponse(w http.ResponseWriter, res interface{}, headers map[string]string, statusCode int) error {
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

	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(statusCode)
	return nil
}
