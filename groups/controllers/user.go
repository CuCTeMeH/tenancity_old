package groups

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userID")
	//db := r.WithContext("db")(gorm.DB)
	render.JSON(w, r, userId) // A chi router helper for serializing and returning json
}
