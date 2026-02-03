package main

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"example.com/chat-server/client"
	"example.com/chat-server/message"
)

const PORT int = 9000

func login(c net.Conn, hubChan chan message.Message, clientRegChan chan client.Client, clientUnregChan chan client.Client) {
	r := bufio.NewScanner(c)
	w := bufio.NewWriter(c)

	fmt.Fprintln(w, "Please enter your name:")
	w.Flush()

	if !r.Scan() {
		return
	}
	name := r.Text()

	fmt.Printf("%s has joined the server \n", name)
	fmt.Fprintf(w, "Welcome %s! You are now connected to the chat server! \n", name)
	w.Flush()

	var user client.Client = client.New(c, make(chan message.Message), name, hubChan)
	clientRegChan <- user
	hubChan <- message.Message{Content: fmt.Sprintf("%s joined the server", name), Kind: message.KindServer}
	user.Activate()

	clientUnregChan <- user
	hubChan <- message.Message{Content: fmt.Sprintf("%s left the server", name), Kind: message.KindServer}
}

func hub(hubChan chan message.Message, clientRegChan chan client.Client, clientUnregChan chan client.Client) {
	clients := make([]client.Client, 0)

	for {
		select {
		case client := <-clientRegChan:
			clients = append(clients, client)
		case client := <-clientUnregChan:
			for i, c := range clients {
				if c.Conn == client.Conn {
					clients = append(clients[:i], clients[i+1:]...)
					close(c.Ch)
					break
				}
			}
		case msg := <-hubChan:
			for _, client := range clients {
				client.Ch <- msg
			}
		}
	}
}

func main() {
	address := fmt.Sprintf(":%v", PORT)

	ln, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening at %s \n", address)

	// create the Hub channel for communicating with all the clients
	hubChan := make(chan message.Message)
	clientRegChan := make(chan client.Client)
	clientUnregChan := make(chan client.Client)
	go hub(hubChan, clientRegChan, clientUnregChan)

	// loop runs for every new connection made
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go login(conn, hubChan, clientRegChan, clientUnregChan)
	}
}
