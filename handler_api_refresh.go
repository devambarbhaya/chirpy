package main

import (
	"chirpy/internal/auth"
	"errors"
	"net/http"
	"strings"
	"time"
)

func (apiCfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	refreshToken := strings.TrimPrefix(authHeader, "Bearer ")
	if refreshToken == "" {
		respondWithError(w, http.StatusUnauthorized, "Missing refresh token", errors.New(""))
		return
	}

	user, err := apiCfg.db.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid refresh token", err)
		return
	}

	accessToken, err := auth.MakeJWT(user.ID, apiCfg.jwtSecret, time.Hour)
	if err != nil {
		respondWithError(w, http. StatusInternalServerError, "Couldn't create access token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: accessToken,
	})
}