package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fardinabir/auth-guard/services/warden/internal/token"
)

type TokenHandler struct {
	tokenAuth *token.TokenAuth
}

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{
		tokenAuth: token.NewTokenAuth(),
	}
}

func (h *TokenHandler) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	tokens := h.tokenAuth.GenerateTokens(req.Username)
	log.Println("Generated tokens successfully:", tokens)
	resp := &token.Response{Status: 200, Body: tokens}
	resp.JSONResponse(w)
}

func (h *TokenHandler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")
	claims, err := token.ValidateToken(tokenStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	log.Println("Validated token successfully:", claims)
	// Set individual headers for each claim
	w.Header().Set("X-Auth-Authorized", fmt.Sprintf("%t", claims.Authorized))
	w.Header().Set("X-Auth-Expiry", fmt.Sprintf("%d", claims.Expiry))
	w.Header().Set("X-Auth-TokenType", claims.TokenType)
	w.Header().Set("X-Auth-Username", claims.UserName)
	resp := &token.Response{Status: 200, Body: claims}
	resp.JSONResponse(w)
}
