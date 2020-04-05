package groups

import (
	"Tenancity/API/groups/controllers"
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{userID}", groups.GetUser)

	return router
}
