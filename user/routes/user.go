package user

import (
	"github.com/go-chi/chi"
	"tenancity/api/user/controllers"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{userID}", user.GetUser)

	return router
}
