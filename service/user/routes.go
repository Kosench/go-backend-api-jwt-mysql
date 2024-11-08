package user

import (
	"github.com/gorilla/mux"
	"go-backend-api-jwt-mysql/types"
	"go-backend-api-jwt-mysql/utils"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// RegisterRoutes registers routes for the server API.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POSt")
	router.HandleFunc("/register", h.handleRegister).Methods("POSt")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	//check if the user exists
	//if it doesnt we createthe new user
}
