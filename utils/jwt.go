package utils

import (
	"errors"
	"time"

	"github.com/ericoliveiras/alert-bot-go/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type Claims struct {
	UserID        uuid.UUID `json:"id"`
	DiscordUserID string    `json:"discord_id"`
	DiscordToken  *oauth2.Token
	jwt.RegisteredClaims
}

var jwtSecret = config.LoadAuthConfig().AccessSecret

func GenerateToken(id uuid.UUID, discordID string, discordToken *oauth2.Token) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)

	claims := Claims{
		UserID:        id,
		DiscordUserID: discordID,
		DiscordToken:  discordToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func GetIDsFromToken(tokenString string) (uuid.UUID, string, *oauth2.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return uuid.Nil, "", nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, claims.DiscordUserID, claims.DiscordToken, nil
	}

	return uuid.Nil, "", nil, errors.New("invalid token")
}
