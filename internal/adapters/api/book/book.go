package book

import (
	"encoding/json"
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/domain"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/gxrlxv/library-rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	bookURL  = "/api/books/:book_id"
	booksURL = "/api/books"
)

type bookHandler struct {
	bookService service.Book
	logger      *logging.Logger
}

func NewBookHandler(service service.Book, logger *logging.Logger) api.Handler {
	return &bookHandler{bookService: service, logger: logger}
}

func (h *bookHandler) Register(router *httprouter.Router) {
	router.POST(booksURL, h.CreateBook)
	router.GET(bookURL, h.GetBook)
	router.GET(booksURL, h.GetAllBooks)
	router.PUT(bookURL, h.UpdateBook)
	router.DELETE(bookURL, h.DeleteBook)
}

func (h *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto domain.CreateBookDTO

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.bookService.CreateBook(r.Context(), dto); err != nil {
		//if err == apperrors.ErrUserNotFound {
		//	http.Error(w, err.Error(), http.StatusNotFound)
		//	return
		//}
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) GetBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *bookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}
