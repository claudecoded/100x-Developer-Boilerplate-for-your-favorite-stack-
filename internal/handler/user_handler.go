package handler

import (
	"encoding/json"
	"net/http"

	"://github.com"
	"://github.com"
	"://github.com"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(router *mux.Router, uc domain.UserUsecase) {
	handler := &UserHandler{usecase: uc}
	
	// Structured REST Routing endpoints Mapping
	router.HandleFunc("/api/v1/users", handler.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users", handler.ListUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/users/{id}", handler.GetUserByID).Methods(http.MethodGet)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload format")
		return
	}

	user, err := h.usecase.Register(r.Context(), input.Name, input.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Malformed UUID query string")
		return
	}

	user, err := h.usecase.GetUser(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.usecase.ListUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, users)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}
