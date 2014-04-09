package datatypes

// The TLV type represents TLV extension fields.
type TLV struct {
	// TODO: Enumeration16
	TlvType		uint16

	// The length of all TLVs shall be an even number of octets.
	LengthField	uint16

	// Octet[lengthField] valueField;
	ValueField	[]byte
}

// The PTPText data type is used to represent textual material in PTP messages.
type PTPText struct {
	LengthField	uint8

	// The textField member shall be encoded as UTF-8 symbols as specified by ISO/IEC 10646:2003. The most significant
	// octet of the leading text symbol shall be the element of the array with index 0.
	// NOTE: A single UTF-8 symbol can be 1–4 octets long. Therefore, the lengthField value can be larger than the
	// number of symbols.
	TextField	[]byte
}

// The FaultRecord type is used to construct fault logs.
type FaultRecord struct {
	// The faultRecordLength member shall indicate the number of octets in the FaultRecord not including the 2 octets
	// of the faultRecordLength member.
	FaultRecordLength 	uint16
	FaultTime			TimeStamp

	// TODO: Enumeration8
	SeverityCode		uint8
	FaultName			PTPText
	FaultValue			PTPText
}


