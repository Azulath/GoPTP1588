package communication

import (
	"net"
	"alex/msc/helper"
)

func MessageSender(ipAddress, port string) {
	address := ipAddress + ":" + port

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	helper.ErrHandling(err)

	connection, err := net.DialUDP("udp", nil, udpAddr)
	helper.ErrHandling(err)
	handleServer(connection)

}

func handleServer(connection *net.UDPConn) {
	//var buffer [512]byte
	_, err := connection.Write([]byte("writening going on"))
	helper.ErrHandling(err)
}

func writePTPHeader(messageType MessageType) []byte {
	//TODO transport specific stuff...
	var header [34]byte
	// transportSpecific AND messageType -> mergen!
	header[0] = messageType
	// reserved AND versionPTP (2)
	header[1] = 2
	// messageLength uint16
	header[2], header[3] = 0, 0
	// domainNumber

	return header
}

