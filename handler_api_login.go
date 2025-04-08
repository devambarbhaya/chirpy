package main

import (
	"chirpy/internal/auth"
	"encoding/json"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email string `json:"email"`
		Password string `json:"password"`
		ExpiresInSeconds int `json:"expires_in_seconds,omitempty"`
	}

	type response struct {
        ID string `json:"id"`
        CreatedAt time.Time `json:"created_at"`
        UpdatedAt time.Time `json:"updated_at"`
        Email string `json:"email"`
        Token string `json:"token"`
    }

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	dbUser, err := apiCfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", nil)
		return
	}

	err = auth.CheckPasswordHash(dbUser.HashedPassword, params.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", nil)
		return
	}

	expiresIn := time.Hour
	if params.ExpiresInSeconds > 0 {
		if params.ExpiresInSeconds > 3600 {
			expiresIn = time.Hour
		} else {
			expiresIn = time.Duration(params.ExpiresInSeconds) * time.Second 
		}
	}

	token, err := auth.MakeJWT(dbUser.ID, apiCfg.jwtSecret, expiresIn)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error creating JWT token", err)
	}

	resp := response{
		ID: dbUser.ID.String(),
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Email: dbUser.Email,
		Token: token,
	}

	respondWithJSON(w, http.StatusOK, resp)
}