package server

type Config struct {
	Addrs []Addr
}

type Server struct {
	cfg *Config
}

// Addr contains string representation of address and its protocol (tcp, unix...).
type Addr struct {
	Proto string
	Addr  string
}

func New(cfg *Config) (*Server, error) {
	s := &Server{
		cfg: cfg,
	}

	return s, nil
}
