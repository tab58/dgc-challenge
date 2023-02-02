package server

type server struct{}

func (s *server) Start() {

}

func NewServer() (*server, error) {
	return &server{}, nil
}
