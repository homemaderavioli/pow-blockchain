package httpserver

import "net/http"

func (s *Server) handleGossip() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//if r.Method != http.MethodPost {
		//	respondWithError(w, r, errors.New("bad request"), http.StatusBadRequest)
		//	return
		//}

		//theirBlockchain := ""
		//theirPeers := ""
		//updateBlockchain(theirBlockchain)
		//updatePeers(theirPeers)

		blocks := s.Blockchain.Blocks()
		respond(w, r, blocks, http.StatusOK)
	}
}
