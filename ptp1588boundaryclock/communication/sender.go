package communication

import (
	"net"
	"alex/msc/helper"
	"alex/ptp1588boundaryclock/datasets"
)

// Global Variables in communication! (concurrency)
var (
	portDS datasets.PortDS
	defaultDS datasets.DefaultDS
	timePropertiesDS datasets.TimePropertiesDS
	msgType MessageType
)

// Prepares and sends the UDP packet to the specified IP address and port
func MessageSender(ipAddress, port string, portDs datasets.PortDS, defaultDs datasets.DefaultDS, timePropertiesDs datasets.TimePropertiesDS, messageType MessageType) {
	address := ipAddress + ":" + port
	portDS = portDs
	defaultDS = defaultDs
	timePropertiesDS = timePropertiesDs
	msgType = messageType

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	helper.ErrHandling(err)

	connection, err := net.DialUDP("udp", nil, udpAddr)
	helper.ErrHandling(err)
	handleServer(connection, writePTPHeader())

}

// Writes 'on' the connection
func handleServer(connection *net.UDPConn, header []byte) {
	//var buffer [512]byte
	//_, err := connection.Write([]byte("writening going on"))
	_, err := connection.Write(header)
	helper.ErrHandling(err)
}
