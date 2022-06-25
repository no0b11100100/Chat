package app

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	"Chat/Client/Server/common"
)

type RemoteServer struct {
	conn   net.Conn
	reader *bufio.Reader
	notify common.ChannelType
	//TODO: use sync.Map
	channels map[common.CommandType][]common.ChannelType
}

func NewRemoteServer(ch common.ChannelType) *RemoteServer {
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

	return &RemoteServer{conn, reader, ch, make(map[common.CommandType][]common.ChannelType)}
}

func (s *RemoteServer) Serve() {
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

		var response common.CommandResponce

		if err := json.Unmarshal(bytes, &response); err != nil {
			log.Println(err)
		}

		if response.Type == common.Notification {
			s.notify <- response.Command
			return
		}

		//Response
		if value, ok := s.channels[response.Command.Type]; ok {
			value[0] <- response.Command
			s.channels[response.Command.Type] = value[1:]
		}

	}
}

func (s *RemoteServer) Send(c common.Command, ch common.ChannelType) {
	bytes, err := json.Marshal(c)

	if err != nil {
		log.Println(err)
	}

	payload := base64.StdEncoding.EncodeToString(bytes)
	payload += "\n"

	if ch != nil {
		s.channels[c.Type] = append(s.channels[c.Type], ch)
	}

	s.conn.Write([]byte(payload))
}
