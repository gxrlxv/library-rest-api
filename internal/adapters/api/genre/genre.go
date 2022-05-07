package genre

import (
	"github.com/gxrlxv/library-rest-api/internal/adapters/api"
	"github.com/gxrlxv/library-rest-api/internal/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	genresURL = "/api/genres"
	genreURL  = "/api/genres/:genre_id"
)

type genreHandler struct {
	genreService service.Genre
}

func NewGenreHandler(service service.Genre) api.Handler {
	return &genreHandler{genreService: service}
}

func (h *genreHandler) Register(router *httprouter.Router) {
	router.POST(genresURL, h.CreateGenre)
	router.GET(genreURL, h.GetGenre)
	router.GET(genresURL, h.GetAllGenres)
	router.PUT(genreURL, h.UpdateGenre)
	router.DELETE(genreURL, h.DeleteGenre)
}

func (h *genreHandler) CreateGenre(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *genreHandler) GetGenre(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *genreHandler) GetAllGenres(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *genreHandler) UpdateGenre(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}

func (h *genreHandler) DeleteGenre(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("implement me")
}
