package book

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
	bookURL      = "/api/books/:book_id"
	booksURL     = "/api/books"
	takeBooksURL = "/api/books/:book_id/take"
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
	router.PUT(takeBooksURL, h.TakeBook)
	router.DELETE(takeBooksURL, h.GiveBook)
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
		if err == apperrors.ErrBookNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) GetBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("book_id")

	book, err := h.bookService.GetBookByID(r.Context(), id)
	if err != nil {
		if err == apperrors.ErrBookNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalBook, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(marshalBook)
	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	books, err := h.bookService.GetAllBooks(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	marshalBooks, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(marshalBooks)
	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto domain.UpdateBookDTO

	id := params.ByName("book_id")

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.bookService.UpdateBook(r.Context(), dto, id)
	if err != nil {
		if err == apperrors.ErrBookNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("book_id")

	err := h.bookService.DeleteBook(r.Context(), id)
	if err != nil {
		if err == apperrors.ErrBookNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) TakeBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var dto domain.TakeBookDTO

	id := params.ByName("book_id")

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.bookService.TakeBook(r.Context(), id, dto.OwnerName)
	if err != nil {
		if err == apperrors.ErrBookNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *bookHandler) GiveBook(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id := params.ByName("book_id")

	err := h.bookService.GiveBook(r.Context(), id)
	if err != nil {
		if err == apperrors.ErrBookNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
