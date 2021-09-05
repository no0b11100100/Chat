package main

import (
	"bufio"
	"command"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

func HandleStartUp(conn net.Conn, serverReader *bufio.Reader) {
	payload, _ := json.Marshal(command.Command{ID: command.StartConnection})
	fmt.Fprintf(conn, string(payload)+"\n")
	_, _ = serverReader.ReadString('\n')
}

func HandleLogIn(conn net.Conn, userReader *bufio.Reader, serverReader *bufio.Reader) {
	const (
		LogIn    = "LogIn"
		Guest    = "Guest"
		Register = "Register"
	)

	fmt.Print("Hello\nWhat action would you do?\nLogIn\nGuest\nRegister\n")
	text, _ := userReader.ReadString('\n')
	fmt.Println("user input", text)
	text = strings.TrimSuffix(text, "\n")

	if strings.EqualFold(LogIn, text) {
		commandPayload := command.UserLoginPayload{}
		fmt.Print("Enter email: ")
		text, _ = userReader.ReadString('\n')
		commandPayload.Email = strings.TrimSuffix(text, "\n")
		fmt.Print("Enter password: ")
		// fmt.Print("\033[8m")
		text, _ = userReader.ReadString('\n')
		// fmt.Print("\033[28m")
		commandPayload.Password = strings.TrimSuffix(text, "\n")
		payload, _ := json.Marshal(commandPayload)
		command := command.Command{ID: command.LogInUser, Payload: payload}
		response, _ := json.Marshal(command)
		fmt.Fprintf(conn, string(response)+"\n")
		// read message from server
		message, _ := serverReader.ReadString('\n')
		fmt.Print("Message from server: " + message)
	} else if strings.EqualFold(text, Guest) {
		command := command.Command{ID: command.GuestUser}
		response, _ := json.Marshal(command)
		fmt.Fprintf(conn, string(response)+"\n")
		// TODO: read message from server
		message, _ := serverReader.ReadString('\n')
		fmt.Print("Message from server: " + message)
	} else if strings.EqualFold(text, Register) {
		commandPayload := command.UserLoginPayload{}
		fmt.Print("Enter email: ")
		text, _ = userReader.ReadString('\n')
		commandPayload.Email = strings.TrimSuffix(text, "\n")
		fmt.Print("Enter password: ")
		password, _ := userReader.ReadString('\n')
		commandPayload.Password = strings.TrimSuffix(password, "\n")
		fmt.Print("Confirm password: ")
		confirwPassword, _ := userReader.ReadString('\n')
		for password != confirwPassword {
			fmt.Print("passwords are not equal, please try again\nConfirm password: ")
			confirwPassword, _ = userReader.ReadString('\n')
		}
		fmt.Print("Enter nick name: ")
		text, _ = userReader.ReadString('\n')
		commandPayload.NickName = strings.TrimSuffix(text, "\n")

		payload, _ := json.Marshal(commandPayload)
		command := command.Command{ID: command.RegisterUser, Payload: payload}
		response, _ := json.Marshal(command)
		fmt.Fprintf(conn, string(response)+"\n")
		// TODO: read message from server
		message, _ := serverReader.ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}

func HandleQuit(conn net.Conn, serverReader *bufio.Reader) {
	command := command.Command{ID: command.Quit}
	response, _ := json.Marshal(command)
	fmt.Fprintf(conn, string(response)+"\n")
	// TODO: read message from server
	message, _ := serverReader.ReadString('\n')
	fmt.Print("Message from server: " + message)
}

func HandleActiveusers(conn net.Conn, serverReader *bufio.Reader) {
	command := command.Command{ID: command.ActiveUsers}
	response, _ := json.Marshal(command)
	fmt.Fprintf(conn, string(response)+"\n")
	// TODO: read responce
	message, _ := serverReader.ReadString('\n')
	fmt.Print("Message from server: " + message)
}

func HandleSendMessage(conn net.Conn, userReader *bufio.Reader, serverReader *bufio.Reader) {
	fmt.Print("Message: ")
	text, _ := userReader.ReadString('\n')
	commandPayload, _ := json.Marshal(command.MessagePayload{Message: text})
	command := command.Command{ID: command.SendMessage, Payload: commandPayload}
	response, _ := json.Marshal(command)
	fmt.Fprintf(conn, string(response)+"\n")
	message, _ := serverReader.ReadString('\n')
	fmt.Print("Message from server: " + message)
}
