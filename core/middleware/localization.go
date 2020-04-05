package core

import (
	"context"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"net/http"
)

func Locale() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			locale := r.URL.Query().Get("locale")

			if locale != "" {
				languageLocale, err := language.Parse(locale)

				if err == nil {
					p := message.NewPrinter(languageLocale)
					r = r.WithContext(context.WithValue(r.Context(), "api.locale", p))
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}
