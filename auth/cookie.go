package auth

import (
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

func SaveCookie(w http.ResponseWriter, token *oauth2.Token) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    token.AccessToken,
		Expires:  expiration,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

func ClearCookie(w http.ResponseWriter) {
	pastTime := time.Now().AddDate(0, 0, -1)
	cookie := http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  pastTime,
		MaxAge:   -1,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}
