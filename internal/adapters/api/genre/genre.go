package genre

import (
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/domain/genre"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	genresURL = "/genres"
	genreURL  = "/genres/:genre_id"
)

type genreHandler struct {
	genreService genre.Service
}

func NewGenreHandler(service genre.Service) api.Handler {
	return &genreHandler{genreService: service}
}

func (h *genreHandler) Register(router *httprouter.Router) {
	router.GET(genresURL, h.GetAllGenres)
}

func (h *genreHandler) GetAllGenres(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	w.Write([]byte("genres"))
	w.WriteHeader(http.StatusOK)
}
