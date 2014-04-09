package communication

import (
	"net"
	"alex/msc/helper"
)

func MessageSender (ipAddress, port string) {
	address := ipAddress + ":" + port

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	helper.ErrHandling(err)

	connection, err := net.DialUDP("udp", nil, udpAddr)
	helper.ErrHandling(err)
	handleServer(connection)

}

func handleServer (connection *net.UDPConn) {
	//var buffer [512]byte
	_, err := connection.Write([]byte("writening going on"))
	helper.ErrHandling(err)
}

