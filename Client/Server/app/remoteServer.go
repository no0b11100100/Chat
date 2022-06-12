package app

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"Chat/Client/Server/channels"
	"Chat/Client/Server/common"
)

type RemoteServer struct {
	conn   net.Conn
	reader *bufio.Reader
}

func NewRemoteServer() *RemoteServer {
	var conn net.Conn
	var err error
	for i := 0; i < 5; i++ {
		conn, err = net.Dial("tcp", "localhost:8081")
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	log.Println("Connected to remote server")

	reader := bufio.NewReader(conn)

	return &RemoteServer{conn, reader}
}

func (s *RemoteServer) Serve(notification *channels.Channels) {
	for {
		payload, err := s.reader.ReadString('\n')
		if err != nil || payload == "" || payload == "\n" {
			continue
		}

		log.Println(string(payload))

		var bytes []byte
		bytes, err = base64.StdEncoding.DecodeString(payload)
		if err != nil {
			log.Println(err)
		}

		var responce common.CommandResponce

		if err := json.Unmarshal(bytes, &responce); err != nil {
			log.Println(err)
		}

		notification.Notify(responce.Command)
	}
}

func (s *RemoteServer) Send(c common.Command) {
	bytes, err := json.Marshal(c)

	if err != nil {
		log.Println(err)
	}

	payload := base64.StdEncoding.EncodeToString(bytes)
	payload += "\n"

	s.conn.Write([]byte(payload))
}
