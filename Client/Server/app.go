package main

import (
	"Chat/Client/Server/communicators"
	"Chat/Client/Server/interfaces"
	"context"
)

type app struct {
	uiServer     interfaces.Server
	remoteServer interfaces.Server
}

func NewApp() *app {
	remoteServer := communicators.NewRemoteServer(":8081")
	return &app{
		uiServer:     communicators.NewGRPCServer(":8080", remoteServer),
		remoteServer: remoteServer,
	}
}

func (a *app) Serve() {
	go a.uiServer.Serve()
	go a.remoteServer.Serve()
	ctx := context.Background()

	<-ctx.Done()
}
