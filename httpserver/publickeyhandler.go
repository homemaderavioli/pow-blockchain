package httpserver

import (
	"errors"
	"net/http"
	pki "pow-blockchain/pki"
)

func (s *Server) handlePublicKey() http.HandlerFunc {
	var response struct {
		PublicKey string `json:"public_key"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		publicKey, err := pki.GetBase64PublicKey(s.NodeKeyPair.PublicKey)
		if err != nil {
			respondWithError(w, r, errors.New("there was a problem"), http.StatusInternalServerError)
			return
		}

		response.PublicKey = publicKey
		respond(w, r, response, http.StatusOK)
	}
}
