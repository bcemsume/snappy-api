package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID, RestaurantID uint
	jwt.StandardClaims
}

var jwtKey = []byte("snappy-my_secret_key")

// CreateJWT s
func CreateJWT(claims jwt.Claims) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {

		return ""
	}
	return tokenString
}

// ValidateJWT s
func ValidateJWT(tknStr string) (bool, interface{}) {

	restClaims := &Claims{}

	t, e := jwt.ParseWithClaims(tknStr, restClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if e != nil {
		return false, nil
	}
	return t.Valid, t.Claims.(*Claims)
}
