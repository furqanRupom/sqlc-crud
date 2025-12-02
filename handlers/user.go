package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/furqanrupom/sqlc-crud/model"
	"github.com/furqanrupom/sqlc-crud/services"
	"github.com/furqanrupom/sqlc-crud/util"
	"github.com/go-chi/chi"
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

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "Params id is missing!",
			Data:    nil,
		})
		return
	}
	id64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "ID is not valid!",
			Data:    err,
		})
		return
	}
	var updateData model.RUser
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&updateData)
	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "Invalid update data",
			Data:    err,
		})
		return
	}
	updateData.ID = int(id64)
	res, err := h.service.Update(r.Context(), updateData)
	if err != nil {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "User update failed",
			Data:    err,
		})
		return
	}
	util.SendResponse(w, http.StatusOK, util.APIResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    res,
	})

}
func(h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request){
	id := chi.URLParam(r, "id")

	if id == "" {
		util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "Params id is missing!",
			Data:    nil,
		})
		return
	}
	err := h.service.Delete(r.Context(),id)
	fmt.Println(err)
	if err != nil {
			util.SendResponse(w, http.StatusBadRequest, util.APIResponse{
			Success: false,
			Message: "User delete failed!",
			Data:    err,
		})
		return
	}
	util.SendResponse(w,http.StatusOK,util.APIResponse{
		Success: true,
		Message: "User successfully deleted!",
		Data:nil,
	})
}
