package communication

import (
	"net"
	"alex/msc/helper"
	"alex/ptp1588boundaryclock/datasets"
	"alex/ptp1588boundaryclock/communication/general"
)

// Prepares and sends the UDP packet to the specified IP address and port
// Signaling and Management are OPTIONAL!
func MessageSender(ipAddress, port string,
	defaultDS datasets.DefaultDS, currentDS datasets.CurrentDS, parentDS datasets.ParentDS,
	portDs datasets.PortDS, timePropertiesDS datasets.TimePropertiesDS,
	msgType MessageType) {
	address := ipAddress + ":" + port

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	helper.ErrHandling(err)

	connection, err := net.DialUDP("udp", nil, udpAddr)
	helper.ErrHandling(err)

	msgSlice := make([]byte, getMsgLength(msgType))
	msgHeader := &Header{&defaultDS, &portDs, &timePropertiesDS, msgType}
	// Message is Interface -> msgText := new(Message) => pointer to interface! bad!
	var msgText Message

	// Maybe function for this...
	// TODO: Rest
	if msgType == Announce {
		msgText = &general.AnnounceMessage{currentDS.StepsRemoved, &parentDS, &timePropertiesDS}
	}

	done := make(chan bool)
	go msgHeader.Write(msgSlice[:34], done)
	go msgText.Write(msgSlice[34:], done)

	if <-done && <-done {
		handleServer(connection, msgSlice)
	}
}

// Writes 'on' the connection
func handleServer(connection *net.UDPConn, msg []byte) {
	//var buffer [512]byte
	//_, err := connection.Write([]byte("writening going on"))
	_, err := connection.Write(msg)
	helper.ErrHandling(err)
}

func getMsgLength(msgType MessageType) uint8 {
	return HeaderLength + msgType.GetLength()
}
