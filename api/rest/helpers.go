package rest

import (
	"log"
	"net/http"
)

// APIServer creates a port for the endpoint.
type APIServer struct {
	listenToPort string
}

// APIFunc is a one-off definition of the HTTP handler declaration for each function.
type APIFunc func(http.ResponseWriter, *http.Request) error

// HTTPHandleFunc decorates the APIFunc into an HTTP handler.
func HTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, reader *http.Request) {
		if err := f(writer, reader); err != nil {
			err := WriteJSON(writer, http.StatusBadRequest, APIError{Error: err.Error()})
			if err != nil {
				log.Fatal("Serving HTTP error failed...")
			}
		}
	}
}

// APIError handles error messages for all API methods.
type APIError struct {
	Error string
}
