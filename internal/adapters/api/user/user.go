package user

import (
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/domain/user"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersURL = "/users"
	userURL  = "/users/:user_id"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(service user.Service) api.Handler {
	return &userHandler{userService: service}
}

func (h *userHandler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetAllUsers)
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.Write([]byte("users"))
	w.WriteHeader(http.StatusOK)
}
