package main

import (
	"fmt"
	"alex/ptp1588boundaryclock/datasets"
	"alex/ptp1588boundaryclock/datatypes"
)

func main () {
	//TODO: POWERUP event (see 9.2.6.2), the execution of this event shall preclude assigning a
	//TODO: clockClass value of 6, 7, 13, or 14.
	defaultDS := new(datasets.DefaultDS)
	//defaultDS.TwoStepFlag = false -> bool is auto false
	defaultDS.NumberPorts = 10
	defaultDS.Priority1 = 1
	defaultDS.Priority2 = 1
	defaultDS.ClockIdentity = datatypes.ClockIdentity{0xFF, 0xFF, 0xA5, 0, 0, 0, 0, 0}
	defaultDS.ClockQuality = datatypes.ClockQuality{248, 0, 0}

	currentDS := new(datasets.CurrentDS)
	currentDS.StepsRemoved = 0

	parentDS := new(datasets.ParentDS)
	parentDS.ObservedParentOffsetScaledLogVariance = 0xFFFF
	parentDS.ObservedParentClockPhaseRateChange = 0x7FFFFFFF
	parentDS.GrandmasterIdentity = defaultDS.ClockIdentity
	parentDS.GrandmasterClockQuality = defaultDS.ClockQuality
	parentDS.GrandmasterPriority1 = defaultDS.Priority1
	parentDS.GrandmasterPriority2 = defaultDS.Priority2

	timePropertiesDS := new(datasets.TimePropertiesDS)
	timePropertiesDS.TimeSource = 0xA0

	portDS := new(datasets.PortDS)
	portDS.PortState = 0x01
	portDS.PeerMeanPathDelay = datatypes.TimeInterval{0}
	// E2E -> 8.2.5.4.4
	portDS.DelayMechanism = 0x01
	portDS.VersionNumber = uint8(2)

	fmt.Println(defaultDS)
	fmt.Println(currentDS)
	fmt.Println(parentDS)
	fmt.Println(timePropertiesDS)
	fmt.Println(portDS)
}

