package auth

import (
	"net/http"
	"time"
)

func SaveTokenCookie(w http.ResponseWriter, tokenString string) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "jwt_token",
		Value:    tokenString,
		Expires:  expiration,
		HttpOnly: false,
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
		HttpOnly: false,
	}
	http.SetCookie(w, &cookie)
}
