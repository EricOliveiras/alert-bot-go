package middleware

import (
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

func IsAuthenticated(r *http.Request) bool {
	cookie, err := r.Cookie("jwt_token")
	if err != nil {
		return false
	}

	token := &oauth2.Token{
		AccessToken: cookie.Value,
	}

	if !token.Valid() {
		return false
	}

	expiryTime := token.Expiry
	return expiryTime.Before(time.Now())
}
