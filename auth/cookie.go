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
