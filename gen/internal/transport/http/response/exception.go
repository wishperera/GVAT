package response

import (
	"encoding/json"
	"fmt"
	"github.com/wishperera/GVAT/gen/internal/pkg/uuid"
	"net/http"
)

type Exception struct {
	Code        int       `json:"code"`
	Description string    `json:"description"`
	TraceID     uuid.UUID `json:"trace_id"`
}

func WriteExceptionResponse(w http.ResponseWriter, e Exception, headers map[string]string, statusCode int) error {
	respByt, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("failed to marshal error response due: %s", err)
	}

	_, err = w.Write(respByt)
	if err != nil {
		return fmt.Errorf("failed to write error response due: %s", err)
	}

	for k, v := range headers {
		w.Header().Set(k, v)
	}

	w.WriteHeader(statusCode)
	return nil
}
