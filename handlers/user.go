package handlers

import (
	"encoding/json"
	"github.com/furqanrupom/sqlc-crud/model"
	"github.com/furqanrupom/sqlc-crud/services"
	"github.com/furqanrupom/sqlc-crud/util"
	"github.com/go-chi/chi"
	"net/http"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var data model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "Invalid data!",
			Data:    err,
		})
		return
	}
	user, err := h.service.Create(r.Context(), data)
	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "User creation failed!",
			Data:    err,
		})
		return
	}
	util.SendResponse(w, http.StatusCreated, util.APIResponse{
		Success: true,
		Message: "User created successfully!",
		Data:    user,
	})

}
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.List(r.Context())
	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "Users fetched failed!",
			Data:    err,
		})
		return
	}
	util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
		Success: true,
		Message: "Users fetched successfully!",
		Data:    users,
	})

}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "Params id is missing!",
			Data:    nil,
		})
		return
	}

	user, err := h.service.Get(r.Context(), id)
	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "User fetched failed!",
			Data:    err,
		})
		return
	}

	util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
		Success: true,
		Message: "Fetched user successfully!",
		Data:    user,
	})

}
