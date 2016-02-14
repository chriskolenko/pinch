package server

type Config struct {
	Logging bool
	Version string
}

type Server struct {
	cfg *Config
}

// Addr contains string representation of address and its protocol (tcp, unix...).
type Addr struct {
	Proto string
	Addr  string
}

func New(cfg *Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Wait(waitChan chan error) {
	waitChan <- nil
}

func (s *Server) Close() {

}
