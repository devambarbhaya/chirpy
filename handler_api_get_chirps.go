package main

import (
	"chirpy/internal/database"
	"net/http"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerGetChirps(w http.ResponseWriter, r *http.Request) {
	authorIDStr := r.URL.Query().Get("author_id")

	var dbChirps []database.Chirp
	var err error

	if authorIDStr != "" {
		authorID, err := uuid.Parse(authorIDStr)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author_id format", err)
			return
		}
		dbChirps, err = apiCfg.db.GetChirpsByAuthorID(r.Context(), authorID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps for author", err)
			return
		}
	} else {
		dbChirps, err = apiCfg.db.GetChirps(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
			return
		}
	}

	chirpList := []Chirp{}
	for _, c := range dbChirps {
		chirpList = append(chirpList, Chirp{
			ID: c.ID,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			Body: c.Body,
			UserID: c.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, chirpList)
}