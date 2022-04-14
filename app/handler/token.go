package handler

import (
	"net/http"

	"github.com/JSainsburyPLC/third-party-token-server/app/context"
	"github.com/JSainsburyPLC/third-party-token-server/db"
)

type TokenResponse struct {
	Token string `json:"token"`
}

func GetToken(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	tokenId := ""
	if val, ok := vals["id"]; ok && len(val) >=1 {
		tokenId = val[0]
	}

	response := ctx.Cache.Get(tokenId)
	if response == nil {
		respondError(w, http.StatusNotFound, "Key not found in cache")
		return
	}

	token := TokenResponse{
		Token: response.Token,
	}

	respondJSON(w, http.StatusOK, token)
}

func PostToken(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	data := db.UserContext{}
	err := ParseBody(r.Body, &data)
	if err != nil  {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
    
	
	err = ctx.Cache.Insert(&data)
	if err != nil {
		respondError(w, http.StatusBadRequest, "Unable to insert key into cache")
		return
	}

	respondJSON(w, http.StatusOK, nil)
}
