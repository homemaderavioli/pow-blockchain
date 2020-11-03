package server

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
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
	s.Router.HandleFunc("/send_money", cors(s.handleSendMoney()))
	s.Router.HandleFunc("/public_key", cors(s.handlePublicKey()))
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

func (s *Server) handlePublicKey() http.HandlerFunc {
	var response struct {
		PublicKey string `json:"public_key"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		publicKey, err := getBase64PublicKey(s.NodeKeyPair.PublicKey)
		if err != nil {
			respondWithError(w, r, errors.New("there was a problem"), http.StatusInternalServerError)
			return
		}

		response.PublicKey = publicKey
		respond(w, r, response, http.StatusOK)
	}
}

func getBase64PublicKey(publicKey *rsa.PublicKey) (string, error) {
	publicKeyData, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyData,
	}

	var keyBuffer bytes.Buffer

	err := pem.Encode(&keyBuffer, publicKeyBlock)
	if err != nil {
		return "", err
	}

	base64PublicKey := base64.StdEncoding.EncodeToString(keyBuffer.Bytes())
	return base64PublicKey, nil
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
