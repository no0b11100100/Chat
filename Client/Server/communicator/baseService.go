package communicator

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/common"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type BaseService struct {
	api.UnimplementedBaseServer
	sender RemoteServerInterface
}

func NewBaseService(sender RemoteServerInterface) *BaseService {
	return &BaseService{sender: sender}
}

func (s *BaseService) SignIn(_ context.Context, userData *api.SignIn) (*api.Result, error) {
	fmt.Printf("SignIn %+v\n", *userData)
	c := common.Command{Type: common.SignIn}
	var err error
	c.Payload, err = json.Marshal(*userData)

	if err != nil {
		log.Println(err)
	}

	ch := make(common.ChannelType)
	s.sender.Send(c, ch)

	response := <-ch
	close(ch)
	log.Println(response)

	var result api.Result

	//TODO: save userID

	result.ResponseStatus = int32(response.Status)

	return &result, nil
}

func (s *BaseService) SignUp(_ context.Context, logIn *api.SignUp) (*api.Result, error) {
	fmt.Printf("SignUp %+v\n", *logIn)
	return nil, nil
}
