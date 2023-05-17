package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenToPort string
}

type APIFunc func(http.ResponseWriter, *http.Request) error

// HTTPHandleFunc decorate the APIFunc into an HTTP handler.
func HTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := f(writer, request); err != nil {
			WriteJSON(writer, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

type APIError struct {
	Error string
}

func WriteJSON(writer http.ResponseWriter, status int, value any) error {
	writer.WriteHeader(status)
	writer.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(writer).Encode(value)
}

// NewApiServer creates a server at a given port.
func NewApiServer(listenToPort string) *APIServer {
	return &APIServer{
		listenToPort: listenToPort,
	}
}

// HandleAccounts handles creating, modifying accounts.
func (as *APIServer) HandleAccounts(writer http.ResponseWriter, reader *http.Request) error {
	if reader.Method == "GET" {
		return as.GetAccount(writer, reader)
	}
	if reader.Method == "POST" {
		return as.CreateAccount(writer, reader)
	}
	if reader.Method == "DELETE" {
		return as.DeleteAccount(writer, reader)
	}
	return fmt.Errorf("method not available, %s", reader.Method)
}

// GetAccount handles getting accounts.
func (as *APIServer) GetAccount(writer http.ResponseWriter, reader *http.Request) error {
	return nil
}

// CreateAccount handles creating new accounts.
func (as *APIServer) CreateAccount(writer http.ResponseWriter, reader *http.Request) error {
	return nil
}

// DeleteAccount handles deleting existing accounts.
func (as *APIServer) DeleteAccount(writer http.ResponseWriter, reader *http.Request) error {
	return nil
}

// Transfer handles transferring assets to different accounts.
func (as *APIServer) Transfer(writer http.ResponseWriter, reader *http.Request) error {
	return nil
}

func (as *APIServer) StartServer() {
	server := mux.NewRouter()
	server.HandleFunc("/accounts", HTTPHandleFunc(as.HandleAccounts))
	log.Println("JSON API server running on port: ", as.listenToPort)
	http.ListenAndServe(as.listenToPort, server)
}
