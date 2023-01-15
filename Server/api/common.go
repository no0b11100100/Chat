package api

type ServerContext struct {
	ConnectionAddress string
}
type disconnectionCallback = func(string)

// enums
type MessageType int

const (
	Request      = 0
	Notification = 1
)

// structs
type MessageData struct {
	Endpoint string      `json:"Endpoint"`
	Topic    string      `json:"Topic"`
	Payload  string      `json:"Payload"`
	Type     MessageType `json:"Type"`
}

type ServerImpl interface {
	Serve()
	Stop()
}

type Server struct {
	servers []ServerImpl
}

func NewServer() *Server {
	return &Server{servers: make([]ServerImpl, 0)}
}

func (s *Server) AddServer(server ServerImpl) {
	s.servers = append(s.servers, server)
}

func (s *Server) Serve() {
	for _, server := range s.servers {
		server.Serve()
	}
}

func (s *Server) Stop() {}
