package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.handleList(w, r)
	case http.MethodPost:
		h.handlePost(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) HandleUser(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	username := pathParts[len(pathParts)-1]

	switch r.Method {
	case http.MethodGet:
		h.handleGet(w, r, username)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handleList(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling list")

	users, err := h.service.List()
	if err != nil {
		log.Printf("Error while listing from service: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Printf("Error encoding response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling post")
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: Validate error of user already existing
	err = h.service.Create(&user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request, username string) {
	log.Println("Handling post")

	user, err := h.service.Get(username)
	if err != nil {
		log.Printf("Error getting user %v: %v", username, err)
		return
	}

	if user == nil {
		log.Printf("Username %v not found", username)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("Error encoding response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
