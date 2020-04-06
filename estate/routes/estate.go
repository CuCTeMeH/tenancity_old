package estate

import (
	"Tenancity/API/estate/controllers"
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{estateID}", estate.GetEstate)

	return router
}
