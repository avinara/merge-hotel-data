package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/merge-hotel-data/controllers"
	httpSwagger "github.com/swaggo/http-swagger"
)

type RouterInterface interface {
	InitRoutes(mergeHotelDataController *controllers.MergeHotelDataController)
	GetMux() *chi.Mux
}

type router struct {
	mux *chi.Mux
}

func NewRouter() RouterInterface {
	mux := chi.NewRouter()
	return &router{
		mux: mux,
	}
}

func (h *router) GetMux() *chi.Mux {
	return h.mux
}

func (h *router) InitRoutes(mergeHotelDataController *controllers.MergeHotelDataController) {

	h.mux.Group(func(r chi.Router) {

		r.Get("/hotels", mergeHotelDataController.GetHotelData)

	})

	h.mux.With(RemoveContextTypeJSON).Get("/swagger/*", httpSwagger.WrapHandler)
}

func RemoveContextTypeJSON(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Content-Type")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
