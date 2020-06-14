package estate

import (
	"github.com/go-chi/chi"
	"tenancity/api/estate/controllers"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{estateID}", estate.GetEstate)

	return router
}
