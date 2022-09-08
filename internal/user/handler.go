package user

import (
	"net/http"

	"github.com/Olexander753/REST_API_WebServer/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

const (
	usersURL = "/users"
	userURL  = "/users/:id"
)

type handler struct {
}

func NewHanler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.POST(usersURL, h.CreateUser)
	router.GET(userURL, h.GetUserByID)
	router.PUT(userURL, h.UpdateUserByID)
	router.PATCH(userURL, h.PartiallyUpdateUserByID)
	router.DELETE(userURL, h.DeleteUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("List users"))
}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("Create user"))
}

func (h *handler) GetUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("User"))
}

func (h *handler) UpdateUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Update user"))

}

func (h *handler) PartiallyUpdateUserByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Partially update user"))
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Delete user"))
}
