package user

import (
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	userSignUpURL  = "api/users/sign-up"
	usersSignInURL = "api/users/sign-in"
	usersURL       = "api/users"
	userURL        = "api/users/:user_id"
)

type userHandler struct {
	userService service.User
}

func NewUserHandler(service service.User) api.Handler {
	return &userHandler{userService: service}
}

func (h *userHandler) Register(router *httprouter.Router) {
	router.POST(userSignUpURL, h.SignUp)
	router.POST(usersSignInURL, h.SignIn)
	router.GET(userURL, h.GetUser)
	router.GET(usersURL, h.GetAllUsers)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *userHandler) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *userHandler) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}
