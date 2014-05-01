package general

import (
	"alex/ptp1588boundaryclock/datasets"
)

type AnnounceMessage struct {
	StepsRemoved uint16
	ParentDS *datasets.ParentDS
	TimePropertiesDS *datasets.TimePropertiesDS
}

// =====================================================================================================================
// Write Announce Message
// =====================================================================================================================

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
	a.setAnnounceGrandmasterClockQuality(announce[14:18])
	// grandmasterPriority2
	announce[18] = a.ParentDS.GrandmasterPriority2
	// grandmasterIdentity
	a.setAnnounceGrandmasterIdentity(announce[19:27])
	// stepsRemoved
	announce[27], announce[28] = uint8(a.StepsRemoved >> 8), uint8(a.StepsRemoved)
	// timeSource
	announce[29] = a.TimePropertiesDS.TimeSource
	done <- true
}

func (a* AnnounceMessage) setAnnounceGrandmasterClockQuality(quality []byte) {
	quality[0] = a.ParentDS.GrandmasterClockQuality.ClockClass
	quality[1] = a.ParentDS.GrandmasterClockQuality.ClockAccuracy
	quality[2] = uint8(a.ParentDS.GrandmasterClockQuality.OffsetScaledVarianceLog >> 8)
	quality[3] = uint8(a.ParentDS.GrandmasterClockQuality.OffsetScaledVarianceLog)
}

func (a* AnnounceMessage) setAnnounceGrandmasterIdentity(identity []byte) {
	for key, value := range a.ParentDS.GrandmasterIdentity {
		identity[key] = value
	}
}

// =====================================================================================================================
// Read Announce Message
// =====================================================================================================================
