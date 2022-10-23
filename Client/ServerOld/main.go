package main

import (
	"Chat/Client/Server/communicator"
)

func main() {

	c := communicator.NewCommunicator()
	c.Serve()
}

/*
grpc server impl - return chan where to write action from UI

tcp client impl - return chan for notification from server

controller

select {
case action := <-communicationFromGRPC
case data := <- communicationFroTCP
}


*/
