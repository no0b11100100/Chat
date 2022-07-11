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
	errors map[common.CommandStatus]string
}

func NewBaseService(sender RemoteServerInterface) *BaseService {
	errors := map[common.CommandStatus]string{
		common.SignInOK:           "Successfully log in",
		common.SignInError:        "Invalid email or password",
		common.SignUpOK:           "You successfully created a new account",
		common.SignUpInvalidEmail: "This email already in use",
	}

	return &BaseService{sender: sender, errors: errors}
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

	err = json.Unmarshal(response.Payload, &result)
	if err != nil {
		log.Println(err)
	}

	result.ErrorMessage = s.errors[response.Status]

	return &result, nil
}

func (s *BaseService) SignUp(_ context.Context, userData *api.SignUp) (*api.Result, error) {
	fmt.Printf("SignUp %+v\n", *userData)

	if userData.Password != userData.ConfirmedPassword {
		return &api.Result{ErrorMessage: "Passwords not match"}, nil
	}

	c := common.Command{Type: common.SignUp}
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

	err = json.Unmarshal(response.Payload, &result)
	if err != nil {
		log.Println(err)
	}

	result.ErrorMessage = s.errors[response.Status]

	return &result, nil
}
