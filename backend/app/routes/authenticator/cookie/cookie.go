package cookie

import (
	"net/http"
	"time"
)

const (
	TokenName = "jwt"
)

func SetCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: http.SameSiteStrictMode,
		Name:     TokenName,
		Value:    token,
		Path:     "/",
	})
}

func DeleteCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		MaxAge:   -1, // delete cookie
		SameSite: http.SameSiteStrictMode,
		Name:     TokenName,
		Value:    "",
		Path:     "/",
	})
}
