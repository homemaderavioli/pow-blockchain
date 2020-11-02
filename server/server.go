package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	blockchain "pow-blockchain/blockchain"
)

type Server struct {
	NodeKeyPair blockchain.KeyPair
	Blockchain  *blockchain.Blockchain
	Port        string
	Router      *http.ServeMux
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func (s *Server) Routes() {
	s.Router.HandleFunc("/gossip", cors(s.handleGossip()))
}

func cors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Origin, Referer")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		h(w, r)
	}
}

func (s *Server) handleGossip() http.HandlerFunc {
	var response struct {
		blockchain blockchain.Blockchain
	}
	return func(w http.ResponseWriter, r *http.Request) {
		//if r.Method != http.MethodPost {
		//	respondWithError(w, r, errors.New("bad request"), http.StatusBadRequest)
		//	return
		//}

		//theirBlockchain := ""
		//theirPeers := ""
		//updateBlockchain(theirBlockchain)
		//updatePeers(theirPeers)

		fmt.Printf("%v\n", *s.Blockchain)

		response.blockchain = *s.Blockchain
		respond(w, r, response, http.StatusOK)
	}
}

func updateBlockchain(theirBlockchain string) {

}

func updatePeers(theirPeers string) {

}

func respondWithError(w http.ResponseWriter, r *http.Request, err error, code int) {
	log.Printf("respond error: %v", err)
	errObj := struct {
		Error string `json:"error"`
	}{Error: err.Error()}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	err = json.NewEncoder(w).Encode(errObj)
	if err != nil {
		log.Printf("respondWithError: %s", err)
	}
}

func respond(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		respondWithError(w, r, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Printf("respond: %s", err)
	}
}

func decode(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return err
	}
	return nil
}
