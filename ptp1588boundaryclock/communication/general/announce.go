package general

import (
	"alex/ptp1588boundaryclock/datasets"
	"alex/ptp1588boundaryclock/datatypes"
)

func WriteAnnounceMessages(currentDS datasets.CurrentDS, parentDS datasets.ParentDS,
	timePropertiesDS datasets.TimePropertiesDS) []byte {
	announce := make([]byte, 30)
	// originTimestamp (Timestamp) 13.5.2.1
	// TODO: implement
	for i := 0; i < 10; i++ {
		announce[i] = 0
	}
	// currentUtcOffset
	announce[10], announce[11] = uint8(timePropertiesDS.CurrentUtcOffset >> 8), uint8(timePropertiesDS.CurrentUtcOffset)
	// reserved
	announce[12] = 0
	// grandmasterPriority1
	announce[13] = parentDS.GrandmasterPriority1
	// grandmasterClockQualtiy
	setAnnounceGrandmasterClockQuality(announce[14:18], &parentDS.GrandmasterClockQuality)
	// grandmasterPriority2
	announce[18] = parentDS.GrandmasterPriority2
	// grandmasterIdentity
	setAnnounceGrandmasterIdentity(announce[19:27], &parentDS.GrandmasterIdentity)
	// stepsRemoved
	announce[27], announce[28] = uint8(currentDS.StepsRemoved >> 8), uint8(currentDS.StepsRemoved)
	// timeSource
	announce[29] = timePropertiesDS.TimeSource
	return announce
}

func setAnnounceGrandmasterClockQuality(quality []byte, clockQuality *datatypes.ClockQuality) {
	quality[0] = clockQuality.ClockClass
	quality[1] = clockQuality.ClockAccuracy
	quality[2], quality[3] = uint8(clockQuality.OffsetScaledVarianceLog >> 8), uint8(clockQuality.OffsetScaledVarianceLog)
}

func setAnnounceGrandmasterIdentity(identity []byte, clockIdentity *datatypes.ClockIdentity) {
	for key, value := range clockIdentity {
		identity[key] = value
	}
}
