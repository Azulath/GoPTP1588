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
	header[2], header[3] = h.calculateMessageLength()
	// domainNumber -> defaultDS.domainNumber
	header[4] = h.DefaultDS.DomainNumber
	// reserved -> 0
	header[5] = 0
	// flagField
	header[6], header[7] = h.setHeaderFlagField()
	// correction field int64 -> seems unimportant
	for i := 8; i < 16; i++ {
		header[i] = 0
	}
	// reserved
	header[16], header[17], header[18], header[19] = 0, 0, 0, 0
	// sourcePortIdentity PortIdentity
	h.setHeaderPortIdentity(header[20:30])
	// sequenceID of the message
	header[30] = 0
	header[31] = 0
	// controlField -> DEPRECATED
	header[32] = h.setHeaderControlField()
	// logMessageInterval
	header[33] = h.setHeaderLogMessageInterval()
	done <- true
}

// Calculates the length of the message
// Will be larger than the header (34)
func (h *Header) calculateMessageLength() (byte, byte) {
	length := uint16(h.MsgType.GetLength() + HeaderLength)
	return uint8(length >> 8), uint8(length)
}

// Sets the header flag filed as specified in 13.3.2.6
func (h *Header) setHeaderFlagField() (flagField1 byte, flagField2 byte) {
	// TODO: i versteh 13.3 Table 20 net
	// Byte 1
	if h.MsgType == Announce || h.MsgType == Sync || h.MsgType == Follow_Up || h.MsgType == Delay_Resp {
		flagField1 += 0
	}

	if (h.MsgType == Sync || h.MsgType == Pdelay_Resp) && h.DefaultDS.TwoStepFlag {
		flagField1 += (1<<1)
	}

	// TODO: if unicast und rest
	// Byte 2
	if h.MsgType == Announce {
		if h.TimePropertiesDS.Leap61 {
			flagField2 += 1
		}
		if h.TimePropertiesDS.Leap59 {
			flagField2 += (1<<1)
		}
		if h.TimePropertiesDS.CurrentUtcOffsetValid {
			flagField2 += (1<<2)
		}
		if h.TimePropertiesDS.PtpTimescale {
			flagField2 += (1<<3)
		}
		if h.TimePropertiesDS.TimeTraceable {
			flagField2 += (1<<4)
		}
		if h.TimePropertiesDS.FrequencyTraceable {
			flagField2 += (1<<5)
		}
	}
	return
}

// Sets the PortIdentity field
// 13.3.2.8
func (h* Header) setHeaderPortIdentity(identity []byte) {
	for key, value := range h.PortDS.PortIdentity.ClockIdentity {
		identity[key] = value
	}
	identity[8], identity[9] = uint8(h.PortDS.PortIdentity.PortNumber >> 8), uint8(h.PortDS.PortIdentity.PortNumber)
}

// Sets the deprecated ControlField
// 13.3.2.10
func (h* Header) setHeaderControlField() (ctrlField byte) {
	if h.MsgType == Sync {
		ctrlField = 0x00
	} else if h.MsgType == Delay_Req {
		ctrlField = 0x01
	} else if h.MsgType == Follow_Up {
		ctrlField = 0x02
	} else if h.MsgType == Delay_Resp {
		ctrlField = 0x03
	} else if h.MsgType == Management {
		ctrlField = 0x04
	} else {
		ctrlField = 0x05
	}
	return
}

func (h* Header) setHeaderLogMessageInterval() (log byte) {
	if h.MsgType == Announce {
		log = uint8(h.PortDS.LogAnnouncedInterval | 0x00)
	} else if h.MsgType == Sync || h.MsgType == Follow_Up {
		// TODO: Multicast
		log = 0x7f
	} else if h.MsgType == Delay_Resp {
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
