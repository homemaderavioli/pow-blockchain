package httpserver

type Peers map[string]Peer

type Peer struct {
	Name      string
	Address   string
	PublicKey string
}

func (s *Server) addPeer(peer Peer) {
	s.Peers[peer.Name] = peer
}
