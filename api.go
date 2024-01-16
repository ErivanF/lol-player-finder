package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, status int, r any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application.json")
	return json.NewEncoder(w).Encode(r)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

type apiServer struct {
	listenAdress string
}

func NewApiServer(listenAdress string) *apiServer {
	return &apiServer{
		listenAdress: listenAdress,
	}
}

func (s *apiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))
	log.Println("Server running on port: ", s.listenAdress)
	http.ListenAndServe(s.listenAdress, router)
}

func (s *apiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}
	return nil
}

func (s *apiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *apiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *apiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
