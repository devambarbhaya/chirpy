package main

import "net/http"

func (apiCfg *apiConfig)handlerAdminReset(w http.ResponseWriter, r *http.Request) {
	if apiCfg.platform != "dev" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	err := apiCfg.db.DeleteAllUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to reset users", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"reset complete"}`))
}