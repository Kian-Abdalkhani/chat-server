package server

import (
	"bufio"
	"fmt"
	"net"
)

// read other client messages and sent to client
func ServerMsg(c net.Conn, msg string) {
	w := bufio.NewWriter(c)

	fmt.Fprintln(w, msg)
	w.Flush()
}
