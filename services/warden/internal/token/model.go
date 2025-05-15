package token

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type TokenDetails struct {
	Authorized bool        `json:"authorized"`
	TokenType  string      `json:"tokenType"`
	UserName   string      `json:"userName"`
	Expiry     interface{} `json:"expiry"`
	jwt.StandardClaims
}

func (t TokenDetails) Valid() error {
	//TODO implement me
	panic("implement me")
}

type Response struct {
	Status int         `json:"status"`
	Body   interface{} `json:"body"`
}

func (r *Response) JSONResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r.Body)
	return
}
