package model

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type User struct {
	gorm.Model
	UserName string `gorm:"unique" json:"userName"`
	MobileNo string `json:"mobileNo"`
	Password string `json:"password,omitempty"`
}

type UserDetails struct {
	ID        uint
	UserName  string `gorm:"unique" json:"userName"`
	MobileNo  string `json:"mobileNo"`
	CreatedAt time.Time
}

type UserLogin struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

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
