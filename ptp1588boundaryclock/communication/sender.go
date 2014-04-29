package communication

import (
	"net"
	"alex/msc/helper"
	"alex/ptp1588boundaryclock/datasets"
	"alex/ptp1588boundaryclock/communication/general"
)

// Prepares and sends the UDP packet to the specified IP address and port
func MessageSender(ipAddress, port string,
	defaultDS datasets.DefaultDS, currentDS datasets.CurrentDS, parentDS datasets.ParentDS,
	portDs datasets.PortDS, timePropertiesDS datasets.TimePropertiesDS,
	msgType MessageType) {
	address := ipAddress + ":" + port

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	helper.ErrHandling(err)

	connection, err := net.DialUDP("udp", nil, udpAddr)
	helper.ErrHandling(err)
	handleServer(connection, append(WritePTPHeader(defaultDS, portDs, timePropertiesDS, msgType),
		general.WriteAnnounceMessages(currentDS, parentDS, timePropertiesDS)...))

}

// Writes 'on' the connection
func handleServer(connection *net.UDPConn, msg []byte) {
	//var buffer [512]byte
	//_, err := connection.Write([]byte("writening going on"))
	_, err := connection.Write(msg)
	helper.ErrHandling(err)
}
