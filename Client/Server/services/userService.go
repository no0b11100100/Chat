package services

import (
	"Chat/Client/Server/api"
	"Chat/Client/Server/common"
	"Chat/Client/Server/interfaces"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
)

type userService struct {
	api.UnimplementedChatServer
	sender interfaces.Sender
	api.UnimplementedUserServer
	errors map[common.CommandStatus]string
}

func NewUserService(sender interfaces.Sender) *userService {
	errors := map[common.CommandStatus]string{
		common.SignInOK:           "Successfully log in",
		common.SignInError:        "Invalid email or password",
		common.SignUpOK:           "You successfully created a new account",
		common.SignUpInvalidEmail: "This email already in use",
	}
	return &userService{sender: sender, errors: errors}
}

func (chat *userService) IsHandledTopic(command common.CommandType) bool {
	return false
}

func (chat *userService) HandleTopic(common.Command) {}

func (s *userService) SignIn(_ context.Context, userData *api.SignIn) (*api.Response, error) {
	fmt.Printf("SignIn %+v\n", *userData)
	c := common.Command{Type: common.SignIn}
	var err error
	c.Payload, err = json.Marshal(*userData)

	if err != nil {
		log.Println(err)
	}

	response := <-s.sender.Send(c)

	log.Println(response)

	var result api.Response

	err = json.Unmarshal(response.Payload, &result)
	if err != nil {
		log.Println(err)
	}

	result.StatusMessage = s.errors[response.Status]

	return &result, nil
}

func (s *userService) SignUp(_ context.Context, userData *api.SignUp) (*api.Response, error) {
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

	response := <-s.sender.Send(c)

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

func (s *userService) EditUser(context.Context, *api.UserData) (*empty.Empty, error) {
	return nil, nil
}
