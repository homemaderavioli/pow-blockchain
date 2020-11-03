package httpclient

import "testing"

func TestGetBlockchain(t *testing.T) {

	getBlockchain("http://localhost:9001/gossip")
}
