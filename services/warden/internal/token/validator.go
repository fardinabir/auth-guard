package token

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateToken(token string) (*TokenDetails, error) {
	if token == "" {
		log.Println("Empty token found")
		return nil, ErrInvalidToken
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnauthorizedReq
		}
		// TODO: shift to viper config
		signSecret := []byte("fardinabir")
		return signSecret, nil
	})
	if err != nil {
		log.Println("Invalid Tokens")
		return nil, ErrInvalidToken
	}
	claims := parsedToken.Claims.(jwt.MapClaims)

	if ok := parsedToken.Valid; !ok {
		log.Println("Invalid Tokens")
		return nil, ErrInvalidToken
	}
	exp := int64(claims["expiry"].(float64))
	if time.Now().After(time.Unix(exp, 0)) {
		log.Println("Tokens Expired")
		return nil, ErrTokenExpired
	}
	tokDetails := &TokenDetails{
		Authorized: claims["authorized"].(bool),
		Expiry:     exp,
		TokenType:  claims["tokenType"].(string),
		UserName:   claims["userName"].(string),
	}
	return tokDetails, nil
}
