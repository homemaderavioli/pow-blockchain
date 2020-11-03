package httpserver

func (s *Server) Routes() {
	s.Router.HandleFunc("/gossip", cors(s.handleGossip()))
	s.Router.HandleFunc("/send_money", cors(s.handleSendMoney()))
	s.Router.HandleFunc("/public_key", cors(s.handlePublicKey()))
}
