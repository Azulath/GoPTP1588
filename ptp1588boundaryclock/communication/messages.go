package communication

type MessageType uint8

type Message interface {
	Write([]byte, chan bool)
}

const (
	Sync		MessageType = iota
	Delay_Req
	Pdelay_Req
	Pdelay_Resp
	_
	_
	_
	_
	Follow_Up
	Delay_Resp
	Pdelay_Resp_Follow_Up
	Announce
	Signaling
	Management
)

func (msg MessageType) GetLength() (length uint8) {
	if msg == Sync || msg == Delay_Req || msg == Follow_Up {
		length = 10
	} else if msg == Pdelay_Req || msg == Pdelay_Resp || msg == Delay_Resp || msg == Pdelay_Resp_Follow_Up {
		length = 20
	} else if msg == Announce {
		length = 30
	} else {
		length = 0
	}
	return
}

