package communication

import (
	"net"
	"alex/msc/helper"
	"alex/ptp1588boundaryclock/datasets"
)

func MessageSender(ipAddress, port string, portDS datasets.PortDS, defaultDS datasets.DefaultDS, timePropertiesDS datasets.TimePropertiesDS, messageType MessageType) {
	address := ipAddress + ":" + port

	udpAddr, err := net.ResolveUDPAddr("udp", address)
	helper.ErrHandling(err)

	connection, err := net.DialUDP("udp", nil, udpAddr)
	helper.ErrHandling(err)
	handleServer(connection, writePTPHeader(portDS, defaultDS, timePropertiesDS, messageType))

}

func handleServer(connection *net.UDPConn, header []byte) {
	//var buffer [512]byte
	//_, err := connection.Write([]byte("writening going on"))
	_, err := connection.Write(header)
	helper.ErrHandling(err)
}

func writePTPHeader(portDS datasets.PortDS, defaultDS datasets.DefaultDS, timePropertiesDS datasets.TimePropertiesDS, msgType MessageType) []byte {
	//TODO transport specific stuff...
	header := make([]byte, 34)
	// transportSpecific AND messageType -> mergen!
	header[0] = byte(msgType)
	// reserved AND versionPTP (2) -> uint8 => no conversion needed
	header[1] = portDS.VersionNumber
	// messageLength uint16
	header[2], header[3] = calculateMessageLength()
	// domainNumber uint8 -> defaultDS.domainNumber
	header[4] = defaultDS.DomainNumber
	// reserved -> 0
	header[5] = 0
	// flag -> mega many much stuff
	header[6], header[7] = setHeaderFlagField(defaultDS, timePropertiesDS, msgType)
	// correction field int64 -> slice (dürft für uns unwichtig sein)
	for i := 8; i < 16; i++{
		header[i] = 0
	}
	// reserved
	header[16], header[17], header[18], header[19] = 0, 0, 0, 0
	// sourcePortIdentity PortIdentity -> slice
	header[20] = 0
	header[29] = 0
	// sequenceID
	header[30] = 0
	header[31] = 0
	// controlField -> DEPRECATED -> if stuff going on soon
	header[32] = 0
	// logMessageInterval -> kummt auf value und so au, mega muh!
	header[33] = 0
	return header
}

// Calculates the length of the message
// Will be larger than the header (34)
func calculateMessageLength() (byte, byte) {
	length := uint16(35)
	return uint8(length >> 8), uint8(length)
}

func setHeaderFlagField(defaultDS datasets.DefaultDS, timePropertiesDS datasets.TimePropertiesDS, msgType MessageType) (flagField1 byte, flagField2 byte) {
	// TODO: i versteh 13.3 Table 20 net
	if msgType == Announce || msgType == Sync || msgType == Follow_Up || msgType == Delay_Resp {
		flagField1 += 0
	}

	if (msgType == Sync || msgType == Pdelay_Resp) && defaultDS.TwoStepFlag {
		flagField1 += (1<<1)
	}

	// TODO: if unicast und rest

	if msgType == Announce {
		if timePropertiesDS.Leap61 {
			flagField2 += 1
		}
		if timePropertiesDS.Leap59 {
			flagField2 += (1<<1)
		}
		if timePropertiesDS.CurrentUtcOffsetValid {
			flagField2 += (1<<2)
		}
		if timePropertiesDS.PtpTimescale {
			flagField2 += (1<<3)
		}
		if timePropertiesDS.TimeTraceable {
			flagField2 += (1<<4)
		}
		if timePropertiesDS.FrequencyTraceable {
			flagField2 += (1<<5	)
		}
	}
	return
}

