package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenAuth struct {
	SignSecretKey []byte
	ExpiryAccess  time.Duration
	ExpiryRefresh time.Duration
}

func NewTokenAuth() *TokenAuth {
	return &TokenAuth{
		// TODO: shift to viper config
		SignSecretKey: []byte("fardinabir"),
		ExpiryAccess:  2,
		ExpiryRefresh: 4,
	}
}

func (t *TokenAuth) GenerateTokens(userName string) Token {
	accToken := t.newToken(userName, t.ExpiryAccess, "access")
	refToken := t.newToken(userName, t.ExpiryRefresh, "refresh")
	return Token{AccessToken: accToken, RefreshToken: refToken}
}

func (t *TokenAuth) newToken(username string, expiry time.Duration, tokenType string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = TokenDetails{
		Authorized: true,
		TokenType:  tokenType,
		UserName:   username,
		Expiry:     time.Now().Add(time.Minute * expiry).Unix(),
	}

	signedToken, err := token.SignedString(t.SignSecretKey)
	if err != nil {
		fmt.Errorf("error while token signing", err.Error())
	}
	return signedToken
}
