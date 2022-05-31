package author

import (
	"encoding/json"
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/gxrlxv/library-rest-api/pkg/apperrors"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	authorsURL = "/api/authors"
	authorURL  = "/api/authors/:author_id"
)

type handler struct {
	authorService service.Author
	logger        *logging.Logger
}

func NewAuthorHandler(service service.Author, logger *logging.Logger) api.Handler {
	return &handler{authorService: service, logger: logger}
}

func (h *handler) Register(router *httprouter.Router) {
	router.POST(authorsURL, h.CreateAuthor)
	router.GET(authorURL, h.GetAuthor)
	router.GET(authorsURL, h.GetAllAuthors)
	router.PUT(authorURL, h.UpdateAuthor)
	router.DELETE(authorURL, h.DeleteAuthor)
}

func (h *handler) CreateAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto domain.CreateAuthorDTO

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.authorService.CreateAuthor(r.Context(), dto); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) GetAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("author_id")

	author, err := h.authorService.GetAuthorByID(r.Context(), id)
	if err != nil {
		if err == apperrors.ErrAuthorNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalUser, err := json.Marshal(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write(marshalUser)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	authors, err := h.authorService.GetAllAuthors(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalAuthors, err := json.Marshal(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(marshalAuthors)
	w.WriteHeader(http.StatusOK)
}

func (h *handler) UpdateAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *handler) DeleteAuthor(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}
