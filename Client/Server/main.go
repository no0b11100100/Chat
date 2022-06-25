package main

import (
	"Chat/Client/Server/communicator"
)

func main() {

	c := communicator.NewCommunicator()
	c.Serve()
}
