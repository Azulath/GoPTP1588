package datatypes

// The PortIdentity type identifies a PTP port.
type PortIdentity struct {
	ClockIdentity     ClockIdentity
	PortNumber        uint16
}

// The PortAddress type represents the protocol address of a PTP port.
type PortAdress struct {
	// TODO: Enumeration16 networkProtocol -> kommt auf die verwendung an
	// The value of the networkProtocol member shall be taken from the networkProtocol enumeration; see 7.4.1.
	NetworkProtocol    uint16

	// The addressLength is the length in octets (bytes) of the address. The range shall be 1 to 16 octets.
	AddressLength       uint16

	// The addressField member holds the protocol address of a port in the format defined by the mapping annex of the
	// protocol as identified by the networkProtocol member. The most significant octet of the addressField is mapped
	// into the octet of the addressField member with index 0. (Octet[addressLength] addressField;)
	AddressField        []byte
}




