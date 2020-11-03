package httpserver

import "net/http"

func (s *Server) handleSendMoney() http.HandlerFunc {
	var request struct {
		Message string `json:"message"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		err := decode(r, &request)
		if err != nil {
			respondWithError(w, r, err, http.StatusBadRequest)
			return
		}

		err = s.Blockchain.AddToChain(s.NodeKeyPair.PrivateKey, []byte(request.Message))
		if err != nil {
			respondWithError(w, r, err, http.StatusBadRequest)
			return
		}
		respond(w, r, nil, http.StatusCreated)
	}
}
