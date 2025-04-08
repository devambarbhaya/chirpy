package main

import (
	"net/http"
	"strings"
)

func (apiCfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		respondWithError(w, http.StatusUnauthorized, "Authorization header is required", nil)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		respondWithError(w, http.StatusUnauthorized, "Authorization header format must be 'Bearer <token>'", nil)
		return
	}
	refreshToken := parts[1]

	err := apiCfg.db.UpdateRevokeToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke the token", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}