package httpserver

import (
	"net/http"

	blockchain "pow-blockchain/blockchain"
	pki "pow-blockchain/pki"
)

type Server struct {
	NodeKeyPair pki.KeyPair
	Blockchain  *blockchain.Blockchain
	Peers       Peers
	Port        string
	Router      *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}
