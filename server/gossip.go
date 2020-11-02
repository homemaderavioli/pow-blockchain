package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type payload struct{}

func GossipWithPeer(peer string) error {
	peerData, err := doRequest(peer)
	if err != nil {
		return err
	}
	defer peerData.Close()

	data := &payload{}

	err = json.NewDecoder(peerData).Decode(data)
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", data)

	return nil
}

func doRequest(address string) (io.ReadCloser, error) {
	req, err := http.NewRequest(http.MethodGet, address, nil)

	if err != nil {
		return nil, err
	}

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return res.Body, nil
}
