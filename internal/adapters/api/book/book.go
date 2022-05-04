package book

import (
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/domain/book"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type bookHandler struct {
	bookService book.Service
}

func NewHandler(service book.Service) api.Handler {
	return &bookHandler{bookService: service}
}

func (h *bookHandler) Register(router *httprouter.Router) {
	router.GET(booksURL, h.GetAllBooks)
}

func (h *bookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)
}
