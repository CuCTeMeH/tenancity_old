package user

import (
	"Tenancity/API/user/controllers"
	"github.com/go-chi/chi"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{userID}", user.GetUser)

	return router
}
