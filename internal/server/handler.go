package server

func (s *Server) Handler(path string) {
	switch path {
	case "/register":
		s.Mux.HandleFunc(path, LogMiddleWare(s.Register))
	case "/login":
		s.Mux.HandleFunc(path, LogMiddleWare(s.Login))
	}
}
