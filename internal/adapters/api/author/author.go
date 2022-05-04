package author

import (
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/domain/author"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	authorsURL = "/authors"
	authorURL  = "/authors/:author_id"
)

type handler struct {
	authorService author.Service
}

func NewAuthorHandler(service author.Service) api.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(authorsURL, h.GetAllAuthors)
}

func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("authors"))
	w.WriteHeader(http.StatusOK)
}
