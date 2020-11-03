package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	blockchain "pow-blockchain/blockchain"
	server "pow-blockchain/httpserver"
	pki "pow-blockchain/pki"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
		log.Printf("port set to %s", port)
	}

	name := os.Getenv("NAME")
	if name == "" {
		return errors.New("NAME must be set")
	}
	log.Printf("name set to %s", name)

	keyPair := pki.GenerateKeyPair()
	bc := &blockchain.Blockchain{}

	//peer := os.Getenv("PEER")
	//if peer == "" {
	// start new blockchain
	bc = blockchain.New(keyPair.PrivateKey)
	//} else {

	//}

	srv := &server.Server{
		NodeKeyPair: keyPair,
		Blockchain:  bc,
		Port:        port,
		Router:      http.NewServeMux(),
	}

	srv.Routes()
	return http.ListenAndServe(":"+port, srv)
}
