package communicator

import (
	"Chat/Client/Server/api"
	"context"
	"fmt"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type Communicator struct {
	api.UnimplementedBaseServer
}

func NewCommunicator() *Communicator {
	return &Communicator{}
}

func (s *Communicator) LogIn(_ context.Context, logIn *api.UserLogIn) (*api.ID, error) {
	fmt.Printf("LogIn %+v\n", *logIn)
	return &api.ID{Id: "id"}, nil
}

func (s *Communicator) Register(_ context.Context, logIn *api.UserLogIn) (*api.ID, error) {
	fmt.Printf("Register %+v\n", *logIn)
	return nil, nil
}

func (s *Communicator) Logout(_ context.Context, id *api.ID) (*empty.Empty, error) {
	fmt.Printf("LogOut %+v\n", *id)
	return nil, nil
}
