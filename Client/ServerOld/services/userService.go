package communicator

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/common"
	"context"
	"encoding/json"
	"fmt"
	"log"

	_ "Chat/Client/Server/interfaces"

	"github.com/golang/protobuf/ptypes/empty"
)

type UserService struct {
	api.UnimplementedUserServer
	sender RemoteServerInterface
	errors map[common.CommandStatus]string
}

func NewUserService(sender RemoteServerInterface) *UserService {
	errors := map[common.CommandStatus]string{
		common.SignInOK:           "Successfully log in",
		common.SignInError:        "Invalid email or password",
		common.SignUpOK:           "You successfully created a new account",
		common.SignUpInvalidEmail: "This email already in use",
	}

	return &UserService{sender: sender, errors: errors}
}

func (s *UserService) SignIn(_ context.Context, userData *api.SignIn) (*api.Response, error) {
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

	var result api.Response

	err = json.Unmarshal(response.Payload, &result)
	if err != nil {
		log.Println(err)
	}

	result.StatusMessage = s.errors[response.Status]

	return &result, nil
}

func (s *UserService) SignUp(_ context.Context, userData *api.SignUp) (*api.Response, error) {
	fmt.Printf("SignUp %+v\n", *userData)

	if userData.Password != userData.ConfirmedPassword {
		return &api.Response{StatusMessage: "Passwords not match"}, nil
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

	var result api.Response

	if response.Status == common.SignUpOK {
		err = json.Unmarshal(response.Payload, &result)
		if err != nil {
			log.Println(err)
		}
	}

	result.StatusMessage = s.errors[response.Status]

	return &result, nil
}

func (s *UserService) EditUser(context.Context, *api.UserData) (*empty.Empty, error) {
	return nil, nil
}
