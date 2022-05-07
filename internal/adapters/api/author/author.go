package author

import (
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	authorsURL = "/api/authors"
	authorURL  = "/api/authors/:author_id"
)

type handler struct {
	authorService service.Author
}

func NewAuthorHandler(service service.Author) api.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(authorsURL, h.CreateAuthor)
	router.GET(authorURL, h.GetAuthor)
	router.GET(authorsURL, h.GetAllAuthors)
	router.PUT(authorURL, h.UpdateAuthor)
	router.DELETE(authorURL, h.DeleteAuthor)
}

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *handler) GetAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}
