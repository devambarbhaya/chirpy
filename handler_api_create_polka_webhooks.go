package main

import (
	"chirpy/internal/auth"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil || apiKey != apiCfg.polkaKey {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized", err)
		return
	}

	webhookReq := parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&webhookReq)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request JSON body", err)
		return
	}

	if webhookReq.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = apiCfg.db.UpdateUserToChirpyRed(r.Context(), webhookReq.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not found", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}