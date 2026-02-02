package client

import (
	"bufio"
	"fmt"
	"net"

	"example.com/chat-server/message"
)

type Client struct {
	Conn    net.Conn
	Ch      chan message.Message
	Name    string
	hubChan chan message.Message
}

func New(conn net.Conn, ch chan message.Message, name string, hubChan chan message.Message) Client {
	return Client{
		Conn:    conn,
		Ch:      ch,
		Name:    name,
		hubChan: hubChan,
	}
}

func (c Client) Activate() {
	defer c.Conn.Close()

	go c.write()
	c.read()
}

// read client's sent messages
func (c Client) read() {
	r := bufio.NewScanner(c.Conn)

	for r.Scan() {
		c.hubChan <- message.Message{
			SenderName: c.Name,
			Content:    r.Text(),
			Kind:       message.KindUser,
		}
	}
}

// read other client messages and sent to client
func (c Client) write() {
	w := bufio.NewWriter(c.Conn)

	for msg := range c.Ch {
		if msg.Kind == message.KindServer {
			fmt.Fprintln(w, msg.Content)
			w.Flush()
		} else {
			text := fmt.Sprintf("%s: %s", msg.SenderName, msg.Content)
			fmt.Fprintln(w, text)
			w.Flush()
		}
	}
}
