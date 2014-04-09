package communication

type MessageType uint8

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

