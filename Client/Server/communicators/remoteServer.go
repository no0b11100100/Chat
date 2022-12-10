package communicators

import (
	"Chat/Client/Server/common"
	log "Chat/Client/Server/common/logger"
	"Chat/Client/Server/interfaces"
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
	"sync"
	"time"
)

type remoteServer struct {
	notificationObserver interfaces.NotificationObserver
	conn                 net.Conn
	reader               *bufio.Reader
	channels             *sync.Map
	mu                   *sync.Mutex
}

func NewRemoteServer(addr string) *remoteServer {
	var conn net.Conn
	var err error
	for i := 0; i < 5; i++ {
		conn, err = net.Dial("tcp", addr)
		if err != nil {
			fmt.Println(err)
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}

	log.Info.Println("Connected to remote server at", addr)

	reader := bufio.NewReader(conn)

	return &remoteServer{
		conn:     conn,
		reader:   reader,
		channels: new(sync.Map),
		mu:       new(sync.Mutex),
	}
}

func (server *remoteServer) Serve() {
	for {
		payload, err := server.reader.ReadString('\n')
		server.mu.Lock()
		log.Info.Println("PAYLOAD", payload)
		if err != nil || payload == "" || payload == "\n" {
			log.Info.Println("SKIP PAYLOAD", payload, err)
			continue
		}

		log.Info.Println(string(payload))

		var bytes []byte
		bytes, err = base64.StdEncoding.DecodeString(payload)
		if err != nil {
			log.Warning.Println(err)
		}

		var response common.CommandResponce

		if err := json.Unmarshal(bytes, &response); err != nil {
			log.Warning.Println(err)
		}

		if response.Type == common.Notification {
			server.notificationObserver.Handle(response.Command)
			log.Info.Println("Receive notification", response)
		} else {
			log.Info.Println("Respose ID", response.Command.ID)
			if value, loaded := server.channels.LoadAndDelete(response.Command.ID); loaded {
				ch, _ := value.(*chan common.Command)
				*ch <- response.Command
				close(*ch)
			} else {
				log.Warning.Println("Unknown command type", response.Command.ID, response.Command.Type)
			}
		}
		server.mu.Unlock()
	}
}

func (server *remoteServer) Send(c common.Command) chan common.Command {
	server.mu.Lock()
	defer server.mu.Unlock()
	id := fmt.Sprintf("%v_%v", c.Type, time.Now().Nanosecond())
	c.ID = id
	bytes, err := json.Marshal(c)

	if err != nil {
		log.Warning.Println(err)
	}

	payload := base64.StdEncoding.EncodeToString(bytes)
	payload += "\n"
	ch := make(chan common.Command, 1)
	server.channels.Store(id, &ch)
	log.Info.Println("Send ID", id)
	server.conn.Write([]byte(payload))

	return ch
}

func (server *remoteServer) AddNotificationObserver(observer interfaces.NotificationObserver) {
	server.notificationObserver = observer
}
