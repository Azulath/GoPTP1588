package datatypes

// The TimeInterval type represents time intervals.
type TimeInterval struct {
	// The scaledNanoseconds member is the time interval expressed in units of nanoseconds and multiplied by 2^+16.
	// Positive or negative time intervals outside the maximum range of this data type shall be encoded as the largest
	// positive and negative values of the data type, respectively.
	// For example, 2.5 ns is expressed as 0000 0000 0002 8000_{16}
	ScaledNanoseconds    int64
}

// The Timestamp type represents a positive time with respect to the epoch.
type TimeStamp struct {
	// TODO: Shift value => uint64 since Go does not support uint48
	// The secondsField member is the integer portion of the timestamp in units of seconds.
	// UInteger48 secondsField;
	SecondsFiled        uint64

	// The nanosecondsField member is the fractional portion of the timestamp in units of nanoseconds.
	// The nanosecondsField member is always less than 10^9.
	// E.g.: +2.000000001 seconds is represented by secondsField = 0000 0000 0002_{16} and
	// nanosecondsField= 0000 0001_{16}
	NanosecondsField    uint32
}

