package general

import (
	"alex/ptp1588boundaryclock/datasets"
	"alex/ptp1588boundaryclock/datatypes"
)

type AnnounceMessage struct {
	CurrentDS *datasets.CurrentDS
	ParentDS *datasets.ParentDS
	TimePropertiesDS *datasets.TimePropertiesDS
}

func (a *AnnounceMessage) Write(announce []byte, done chan bool) {
	// originTimestamp (Timestamp) 13.5.2.1
	// TODO: implement
	for i := 0; i < 10; i++ {
		announce[i] = 0
	}
	// currentUtcOffset
	announce[10], announce[11] = uint8(a.TimePropertiesDS.CurrentUtcOffset >> 8), uint8(a.TimePropertiesDS.CurrentUtcOffset)
	// reserved
	announce[12] = 0
	// grandmasterPriority1
	announce[13] = a.ParentDS.GrandmasterPriority1
	// grandmasterClockQualtiy
	setAnnounceGrandmasterClockQuality(announce[14:18], &a.ParentDS.GrandmasterClockQuality)
	// grandmasterPriority2
	announce[18] = a.ParentDS.GrandmasterPriority2
	// grandmasterIdentity
	setAnnounceGrandmasterIdentity(announce[19:27], &a.ParentDS.GrandmasterIdentity)
	// stepsRemoved
	announce[27], announce[28] = uint8(a.CurrentDS.StepsRemoved >> 8), uint8(a.CurrentDS.StepsRemoved)
	// timeSource
	announce[29] = a.TimePropertiesDS.TimeSource
	done <- true
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
