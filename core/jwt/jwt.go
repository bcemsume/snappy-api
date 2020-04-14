package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type RestaurantClaims struct {
	RestaurantID uint
	jwt.StandardClaims
}

type UserClaims struct {
	UserID uint
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
func ValidateJWT(tknStr string) bool {

	userClaims := &UserClaims{}
	tkn, err := jwt.ParseWithClaims(tknStr, userClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		restClaims := &RestaurantClaims{}

		t, e := jwt.ParseWithClaims(tknStr, restClaims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if e != nil {
			return false
		}
		return t.Valid
	}

	return tkn.Valid
}
