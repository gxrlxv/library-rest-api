package book

import (
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	bookURL  = "api/books/:book_id"
	booksURL = "api/books"
)

type bookHandler struct {
	bookService service.Book
}

func NewHandler(service service.Book) api.Handler {
	return &bookHandler{bookService: service}
}

func (h *bookHandler) Register(router *httprouter.Router) {
	router.POST(booksURL, h.CreateBook)
	router.GET(bookURL, h.GetBook)
	router.GET(booksURL, h.GetAllBooks)
	router.PUT(bookURL, h.UpdateBook)
	router.DELETE(bookURL, h.DeleteBook)
}

func (h *bookHandler) GetBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *bookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}
func (h *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}
