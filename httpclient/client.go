package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type payload struct{}

type blocks []block

type block struct {
	BlockNumber       int     "json:`BlockNumber`"
	Nonce             int     "json:`Nonce`"
	Message           message "json:`Message`"
	PreviousBLockHash string  "json`PreviousBlockHash`"
}

type message struct {
	Message          string "json:`Message`"
	MessageHash      string "json:`MessageHash`"
	MessageSignature string "json:`MessageSignature`"
}

func getBlockchain(peer string) error {
	peerData, err := doRequest(peer)
	if err != nil {
		return err
	}
	defer peerData.Close()

	data := &blocks{}

	err = json.NewDecoder(peerData).Decode(data)
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", data)

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
