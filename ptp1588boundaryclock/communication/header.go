package communication

import "alex/ptp1588boundaryclock/datasets"

type Header struct {
	DefaultDS *datasets.DefaultDS
	PortDS *datasets.PortDS
	TimePropertiesDS *datasets.TimePropertiesDS
	MsgType MessageType
}

const (
	HeaderLength uint8 = 34
)

// =====================================================================================================================
// Write Header
// =====================================================================================================================

// Creates the header as specified in 13.3
func (h *Header) Write(header []byte, done chan bool) {
	//TODO transport specific stuff...
	// transportSpecific AND messageType
	header[0] = byte(h.MsgType)
	// reserved AND versionPTP (2)
	header[1] = h.PortDS.VersionNumber
	// messageLength
	header[2], header[3] = calculateMessageLength()
	// domainNumber -> defaultDS.domainNumber
	header[4] = h.DefaultDS.DomainNumber
	// reserved -> 0
	header[5] = 0
	// flagField
	header[6], header[7] = setHeaderFlagField(h.DefaultDS, h.TimePropertiesDS, h.MsgType)
	// correction field int64 -> seems unimportant
	for i := 8; i < 16; i++ {
		header[i] = 0
	}
	// reserved
	header[16], header[17], header[18], header[19] = 0, 0, 0, 0
	// sourcePortIdentity PortIdentity
	setHeaderPortIdentity(header[20:30], h.PortDS)
	// sequenceID of the message
	header[30] = 0
	header[31] = 0
	// controlField -> DEPRECATED
	header[32] = setHeaderControlField(h.MsgType)
	// logMessageInterval
	header[33] = setHeaderLogMessageInterval(h.PortDS, h.MsgType)
	done <- true
}

// Calculates the length of the message
// Will be larger than the header (34)
func calculateMessageLength() (byte, byte) {
	length := uint16(35)
	return uint8(length >> 8), uint8(length)
}

// Sets the header flag filed as specified in 13.3.2.6
func setHeaderFlagField(defaultDS *datasets.DefaultDS,
	timePropertiesDS *datasets.TimePropertiesDS, msgType MessageType) (flagField1 byte, flagField2 byte) {
	// TODO: i versteh 13.3 Table 20 net
	// Byte 1
	if msgType == Announce || msgType == Sync || msgType == Follow_Up || msgType == Delay_Resp {
		flagField1 += 0
	}

	if (msgType == Sync || msgType == Pdelay_Resp) && defaultDS.TwoStepFlag {
		flagField1 += (1<<1)
	}

	// TODO: if unicast und rest
	// Byte 2
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
			flagField2 += (1<<5)
		}
	}
	return
}

// Sets the PortIdentity field
// 13.3.2.8
func setHeaderPortIdentity(identity []byte, portDS *datasets.PortDS) {
	for key, value := range portDS.PortIdentity.ClockIdentity {
		identity[key] = value
	}
	identity[8], identity[9] = uint8(portDS.PortIdentity.PortNumber >> 8), uint8(portDS.PortIdentity.PortNumber)
}

// Sets the deprecated ControlField
// 13.3.2.10
func setHeaderControlField(msgType MessageType) (ctrlField byte) {
	if msgType == Sync {
		ctrlField = 0x00
	} else if msgType == Delay_Req {
		ctrlField = 0x01
	} else if msgType == Follow_Up {
		ctrlField = 0x02
	} else if msgType == Delay_Resp {
		ctrlField = 0x03
	} else if msgType == Management {
		ctrlField = 0x04
	} else {
		ctrlField = 0x05
	}
	return
}

func setHeaderLogMessageInterval(portDS *datasets.PortDS, msgType MessageType) (log byte) {
	if msgType == Announce {
		log = uint8(portDS.LogAnnouncedInterval | 0x00)
	} else if msgType == Sync || msgType == Follow_Up {
		// TODO: Multicast
		log = 0x7f
	} else if msgType == Delay_Resp {
		// TODO: Multicast
		log = 0x7f
	} else {
		log = 0x7f
	}
	return
}

// =====================================================================================================================
// Read Header
// =====================================================================================================================
