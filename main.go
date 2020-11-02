package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	blockchain "pow-blockchain/blockchain"
	server "pow-blockchain/server"
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

	//peers := strings.Split(os.Getenv("PEERS"), ",")

	keyPair := blockchain.GenerateKeyPair()
	bc := blockchain.New(keyPair)

	srv := &server.Server{
		NodeKeyPair: keyPair,
		Blockchain:  bc,
		Port:        port,
		Router:      http.NewServeMux(),
	}
	srv.Routes()
	return http.ListenAndServe(":"+port, srv)
}
