package datatypes

// The ClockIdentity type identifies a clock
// In our implementation we merge two octets to one byte. Shifting will be necessary again.
// github test
type ClockIdentity [8]byte

// The ClockQuality represents the quality of a clock.
type ClockQuality struct {
	ClockClass				uint8

	// TODO: Enumeration8 => uint8?
	ClockAccuracy 			uint8
	OffsetScaledVarianceLog	uint16
}
