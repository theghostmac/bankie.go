package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/theghostmac/bankie.go/common/logger"
	"github.com/theghostmac/bankie.go/database/users"
	"net/http"
	"strconv"
)

// WriteJSON writes JSON response to the http.ResponseWriter.
func WriteJSON(writer http.ResponseWriter, status int, value interface{}) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(value)
}

// NewAPIServer creates a new APIServer instance listening on the provided port.
func NewAPIServer(ListenToPort string, Store users.Storage) *APIServer {
	return &APIServer{
		ListenToPort: ListenToPort,
		Store:        Store,
	}
}

// HandleAccounts handles requests for creating, modifying, and deleting accounts.
func (as *APIServer) HandleAccounts(writer http.ResponseWriter, request *http.Request) error {
	if request.Method == "GET" {
		return as.GetAccountByID(writer, request)
	}
	if request.Method == "POST" {
		return as.CreateAccount(writer, request)
	}
	//if request.Method == "DELETE" {
	//	return as.DeleteAccount(writer, request)
	//}
	return fmt.Errorf("method not available: %s", request.Method)
}

// GetAccount gets all accounts.
func (as *APIServer) GetAccount(writer http.ResponseWriter, request *http.Request) error {
	accounts, err := as.Store.GetAccounts()
	if err != nil {
		return err
	}
	return WriteJSON(writer, http.StatusOK, accounts)
}

// GetAccountByID handles GET requests for retrieving account information with ID.
func (as *APIServer) GetAccountByID(writer http.ResponseWriter, request *http.Request) error {
	if request.Method == "GET" {
		id, err := GetID(request)
		if err != nil {
			return err
		}

		account, err := as.Store.GetAccountByID(id)
		if err != nil {
			return err
		}
		return WriteJSON(writer, http.StatusOK, account)
	}
	if request.Method == "DELETE" {
		return as.DeleteAccount(writer, request)
	}
	return fmt.Errorf("method not allowed %s: ", request.Method)
}

func GetID(request *http.Request) (int, error) {
	idVariable := mux.Vars(request)["id"]
	id, err := strconv.Atoi(idVariable)
	if err != nil {
		return id, fmt.Errorf("invalid ID given: %s ", idVariable)
	}
	return id, nil
}

// CreateAccount handles POST requests for creating new accounts.
func (as *APIServer) CreateAccount(writer http.ResponseWriter, request *http.Request) error {
	createAccountRequest := new(users.CreateCustomerRequest)
	if err := json.NewDecoder(request.Body).Decode(createAccountRequest); err != nil {
		return err
	}
	account := users.NewCustomer(
		createAccountRequest.FirstName, createAccountRequest.LastName, createAccountRequest.Email,
	)
	if err := as.Store.CreateAccount(account); err != nil {
		return err
	}
	return WriteJSON(writer, http.StatusOK, createAccountRequest)
}

// DeleteAccount handles DELETE requests for deleting existing accounts.
func (as *APIServer) DeleteAccount(writer http.ResponseWriter, request *http.Request) error {
	id, err := GetID(request)
	if err != nil {
		return err
	}
	if err := as.Store.DeleteAccount(id); err != nil {
		return err
	}
	return WriteJSON(writer, http.StatusOK, map[string]int{"deleted": id})
}

// Transfer handles transferring assets between accounts.
func (as *APIServer) Transfer(writer http.ResponseWriter, request *http.Request) error {
	transferRequest := new(users.TransferRequest)
	if err := json.NewDecoder(request.Body).Decode(transferRequest); err != nil {
		return err
	}
	defer request.Body.Close()
	return WriteJSON(writer, http.StatusOK, transferRequest)
}

// StartServer starts the API server.
func (as *APIServer) StartServer() {
	server := mux.NewRouter()
	server.HandleFunc("/", HTTPHandleFunc(nil))
	server.HandleFunc("/accounts", HTTPHandleFunc(as.HandleAccounts))
	server.HandleFunc("/accounts/{id}", WithJWTAuth(HTTPHandleFunc(as.GetAccountByID)))

	server.HandleFunc("/transfer/", HTTPHandleFunc(as.Transfer))
	//log.Println("JSON API server running on port:", as.listenToPort)
	logger.InfoLogs("JSON API server running on specified port...\n")
	http.ListenAndServe(as.ListenToPort, server)
}
