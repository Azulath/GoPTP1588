package communication

import (
	"net"
	"alex/msc/helper"
)

func MessageListener (port string) {
	port = ":" + port
	udpAddress, err := net.ResolveUDPAddr("udp4", port)
	helper.ErrHandling(err)

	connection, err := net.ListenUDP("udp", udpAddress)
	helper.ErrHandling(err)

	for {
		// infinite loop -> go routine
		handleClient(connection)
	}
}

// TODO: Concurrency
func handleClient (connection *net.UDPConn) {
	var buffer [512] byte
	_, _, err := connection.ReadFromUDP(buffer[0:])
	if err != nil {
		return
	}
}
