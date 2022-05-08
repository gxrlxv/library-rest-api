package user

import (
	"encoding/json"
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	userSignUpURL  = "/api/users/sign-up"
	usersSignInURL = "/api/users/sign-in"
	usersURL       = "/api/users"
	userURL        = "/api/users/:user_id"
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
	var rq CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&rq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.userService.CreateUser(r.Context(), domain.User{Email: rq.Email, Username: rq.Username, PasswordHash: rq.PasswordHash})
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *userHandler) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("user sign in"))
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("get user by id"))
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("get all users"))
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("delete user"))
}
