package communicator

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/common"
	"context"
	"encoding/json"
	"fmt"
	"log"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type BaseService struct {
	api.UnimplementedBaseServer
	sender   RemoteServerInterface
	baseChan <-chan string
}

func NewBaseService(sender RemoteServerInterface, ch <-chan string) *BaseService {
	return &BaseService{sender: sender, baseChan: ch}
}

func (s *BaseService) LogIn(_ context.Context, logIn *api.UserLogIn) (*api.ID, error) {
	fmt.Printf("LogIn %+v\n", *logIn)
	c := common.Command{Type: common.LogIn}
	var err error
	c.Payload, err = json.Marshal(*logIn)

	if err != nil {
		log.Println(err)
	}

	s.sender.Send(c)

	responce := <-s.baseChan
	log.Println(responce)

	var id api.ID
	if err := json.Unmarshal([]byte(responce), &id); err != nil {
		return nil, err
	}

	return &api.ID{Id: id.Id}, nil
}

func (s *BaseService) Register(_ context.Context, logIn *api.UserLogIn) (*api.ID, error) {
	fmt.Printf("Register %+v\n", *logIn)
	return nil, nil
}

func (s *BaseService) Logout(_ context.Context, id *api.ID) (*empty.Empty, error) {
	fmt.Printf("LogOut %+v\n", *id)
	return nil, nil
}
