package user

import (
	"encoding/json"
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/gxrlxv/library-rest-api/pkg/apperrors"
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
	router.PUT(userURL, h.UpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}

func (h *userHandler) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto domain.CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.CreateUser(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *userHandler) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto domain.SignInUserDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.userService.SignIn(r.Context(), dto); err != nil {
		if err == apperrors.ErrUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("user_id")

	user, err := h.userService.GetUserByID(r.Context(), id)
	if err != nil {
		if err == apperrors.ErrUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalUser, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(marshalUser)
	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users, err := h.userService.GetAllUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	marshalUsers, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(marshalUsers)
	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto domain.UpdateUserDTO

	id := params.ByName("user_id")

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.userService.UpdateUser(r.Context(), dto, id)
	if err != nil {
		if err == apperrors.ErrUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *userHandler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("user_id")

	err := h.userService.DeleteUser(r.Context(), id)
	if err != nil {
		if err == apperrors.ErrUserNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
