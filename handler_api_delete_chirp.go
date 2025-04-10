package main

import (
	"chirpy/internal/auth"
	"database/sql"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerDeleteChirp(w http.ResponseWriter, r *http.Request) {
	chirpIDStr := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirpID format", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Missing or malformed access token", err)
		return
	}

	userID, err := auth.ValidateJWT(token, apiCfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid token", err)
		return
	}

	chirp, err := apiCfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		} else {
			respondWithError(w, http.StatusInternalServerError, "Error fetching chirp", err)
		}
		return
	}

	if chirp.UserID != userID {
		respondWithError(w, http.StatusForbidden, "You are not the author of this chirp", err)
		return
	}

	err = apiCfg.db.DeleteChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to delete chirp", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}