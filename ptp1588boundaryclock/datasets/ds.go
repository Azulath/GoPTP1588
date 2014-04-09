package datasets

import "alex/ptp1588boundaryclock/datatypes"

// =====================================================================================================================
// 8.1.2 Initialization Configuration - test
//
// General: Every member of a data set is classified as static, dynamic, or configurable.
//
// Static:
//	- The values of static members are inherent physical or operational properties of the clock or of the protocol.
// Dynamic:
//	- The values of dynamic members are not directly changed by users but may change as follows:
//		-- As a result of protocol operations.
//		-- Due to changes in the internal properties of the clock.
//		-- Due to interactions with timing systems external to PTP.
// Configurable:
//	- The values of configurable members can only be changed using management messages or implementation- specific
// 		configuration means.
//	- Unless otherwise stated in this standard, when the value of a configurable member is updated,
// 		the update value shall take effect immediately upon update.
//	- Unless otherwise stated in this standard, the update values for configurable members shall be restricted
// 		in range to the most restrictive of the range values specified in Clause 7
// 		or specified in the applicable PTP profile.
//
// 8.1.3.1 General initialization specifications - Data set members shall be initialized before leaving the
// 	INITIALIZATION state of an ordinary or boundary clock
//
// 8.1.3.2 Initialization of static data set members
//	- Static members shall be initialized to the implementation-specific value meeting the specifications
// 		for the member.
// 8.1.3.3 Initialization of dynamic data set members
//	- Dynamic members shall be initialized to the first of the following values that applies:
//		-- The value mandated in the specification for the data set member
//		-- The value that represents the properties of the clock or protocol at the time of initialization
//		-- The value in nonvolatile read−write storage if implemented
//		-- Implementation-specific value
// =====================================================================================================================

type DefaultDS struct {
	// =================================================================================================================
	// Static Members
	// =================================================================================================================

	// The value of defaultDS.twoStepFlag shall be TRUE if the clock is a two-step clock.
	TwoStepFlag                                bool

	// The value of defaultDS.clockIdentity shall be the clockIdentity (see 7.6.2.1) of the local clock.
	ClockIdentity                              datatypes.ClockIdentity

	// The value of defaultDS.numberPorts shall be the number of PTP ports on the device.
	// For an ordinary clock, this value shall be 1.
	NumberPorts                                uint16

	// =================================================================================================================
	// Dynamic Members
	// =================================================================================================================
	ClockQuality                            datatypes.ClockQuality

	// =================================================================================================================
	// Configurable Members
	// =================================================================================================================

	// The value of defaultDS.priority1 is the priority1 attribute (see 7.6.2.2) of the local clock.
	Priority1                                uint8

	// The value of defaultDS.priority2 is the priority2 attribute (see 7.6.2.3) of the local clock.
	Priority2                                uint8

	// The value of defaultDS.domainNumber is the domain attribute (see 7.1) of the local clock.
	DomainNumber                             uint8

	// The value of defaultDS.slaveOnly shall be TRUE if the clock is a slave-only clock; see 9.2.2.
	// The value shall be FALSE if the clock is a non-slave-only clock; see 9.2.3.
	SlaveOnly                                bool
}

type CurrentDS struct {
	// =================================================================================================================
	// Dynamic Members
	// =================================================================================================================

	// The value of currentDS.stepsRemoved is the number of communication paths traversed between the local clock
	// and the grandmaster clock.
	// The initialization value shall be 0.
	StepsRemoved                             uint16

	OffsetFromMaster                         datatypes.TimeInterval
	MeanPathDelay                            datatypes.TimeInterval
}

type ParentDS struct {
	// =================================================================================================================
	// Dynamic Members
	// =================================================================================================================

	// The value of parentDS.parentPortIdentity is the portIdentity of the port on the master
	// that issues the Sync messages used in synchronizing this clock.
	ParentPortIdentity                         datatypes.PortIdentity

	// The value of parentDS.parentStats shall be TRUE if all of the following conditions are satisfied:
	//	- The clock has a port in the SLAVE state.
	//	- The clock has computed statistically valid estimates parentDS.observedParentOffsetScaledLog Variance
	// 		and parentDS.observedParentClockPhaseChangeRate members.
	//	- Otherwise the value shall be FALSE.
	// The initialization value shall be FALSE.
	ParentStats                                bool

	ObservedParentOffsetScaledLogVariance      uint16 // !! COMPUTATION OPTIONAL -> IF NOT ParentSTATS FALSE !!

	// The initialization value shall be 7FFF FFFF16
	// irrespective of whether the computation is implemented in the local clock.
	// A value equal to 7FFF FFFF16 indicates that either the value exceeds the capacity of the data type
	// or that the value has not been computed.
	ObservedParentClockPhaseRateChange        int32 // !! COMPUTATION OPTIONAL -> IF NOT ParentSTATS FALSE !!

	// The value of parentDS.grandmasterIdentity is the clockIdentity attribute (see 7.6.2.1) of the grandmaster clock.
	// The initialization value shall be the defaultDS.clockIdentity member.
	GrandmasterIdentity                        datatypes.ClockIdentity

	// The value of parentDS.grandmasterClockQuality is the clockQuality attribute (see 7.6.2.4, 7.6.2.5, and 7.6.3)
	// of the grandmaster clock.
	// The initialization value shall be the value of the defaultDS.clockQuality member.
	GrandmasterClockQuality                    datatypes.ClockQuality

	// The initialization value shall be the value of the defaultDS.priority1 member.
	GrandmasterPriority1                       uint8

	// The initialization value shall be the value of the parentDS.priority2 member.
	GrandmasterPriority2                       uint8
}

type TimePropertiesDS struct {
	// =================================================================================================================
	// Dynamic Members
	// =================================================================================================================

	// Offset between TAI and UTC; otherwise the value has no meaning. The value shall be in units of seconds.
	// -> OPTIONAL?
	CurrentUtcOffset                          int16

	// The initialization value shall be TRUE if the value of timePropertiesDS.currentUtcOffset is known to be correct;
	// otherwise, it shall be FALSE.
	CurrentUtcOffsetValid                     bool

	// In PTP systems whose epoch is the PTP epoch, a TRUE value shall indicate that the last minute of the current
	// UTC day contains 59 seconds.
	// If the epoch is not PTP, the value shall be set to FALSE
	// Initialization:
	//	- If the timePropertiesDS.ptpTimescale (see 8.2.4.8) is TRUE, the value shall be the value obtained
	//		from a primary reference if known at the time of initialization, else.
	//	- The value shall be FALSE
	Leap59                                    bool

	// Like Leap59 - Exception: 61 seconds
	Leap61                                    bool

	// Initialization:
	//	- If the timePropertiesDS.ptpTimescale (see 8.2.4.8) is TRUE and the time and the value of
	// 		timePropertiesDS.currentUtcOffset are traceable to a primary reference at the time of initialization,
	// 		the value shall be TRUE, else
	//	- The value shall be FALSE.
	TimeTraceable                             bool

	// Initialization:
	//	- If the frequency is traceable to a primary reference at the time of initialization the value shall be TRUE
	//	- Else the value shall be FALSE.
	FrequencyTraceable                        bool

	// !! INITIALIZE FIRST !!
	// Initialization:
	//	- If the clock timescale (see 7.2.1) is PTP and this is known at the time of initialization, the value shall
	// 		be set to TRUE, else
	//	- The value shall be FALSE, indicating that the timescale is ARB.
	PtpTimescale                              bool

	// TODO: Enumeration8
	// The value of timePropertiesDS.timeSource is the source of time used by the grandmaster clock.
	// Initialization:
	//	- If the timeSource (see 7.6.2.6) is known at the time of initialization, the value shall be set to that value.
	//	- Else the value shall be INTERNAL_OSCILLATOR.
	TimeSource                                uint8
}

// The number of such data sets shall be the value of defaultDS.numberPorts
type PortDS struct {
	// =================================================================================================================
	// Static Members
	// =================================================================================================================

	// The value of portDS.portIdentity shall be the PortIdentity attribute of the local port; see 7.5.2.
	PortIdentity                            datatypes.PortIdentity

	// =================================================================================================================
	// Dynamic Members
	// =================================================================================================================

	// TODO: Enumeration8
	// The value of portDS.portState shall be the value of the current state of the protocol engine associated with
	// this port (see 9.2) and shall be taken from the enumeration in Table 8.
	//		PTP state enumeration		Value (hex)
	//		INITIALIZING				01
	//		FAULTY						02
	//		DISABLED					03
	//		PRE_MASTER					04
	//		MASTER						06
	//		PASSIVE						07
	//		UNCALIBRATED				08
	//		SLAVE						09
	//		-							All other values reserved
	// The initialization value shall be INITIALIZING.
	PortState                                 uint8

	// The value of portDS.logMinDelayReqInterval is the logarithm to the base 2 of the minDelayReqInterval;
	// see 7.7.2.4. The initialization value is implementation-specific consistent with 7.7.2.4.
	LogMinDelayReqInterval                    int8

	// ? OPTIONAL ?
	// If the value of the portDS.delayMechanism member is peer-to-peer (P2P), the value of portDS.peerMeanPathDelay
	// shall be an estimate of the current one-way propagation delay on the link, i.e., <meanPathDelay>, attached to
	// this port computed using the peer delay mechanism; see 11.4. The data type should be TimeInterval. If the value
	// of the portDS.delayMechanism member is end-to-end (E2E), this member’s value shall be zero.
	// The initialization value shall be zero.
	PeerMeanPathDelay                         datatypes.TimeInterval

	// =================================================================================================================
	// Configurable Members
	// =================================================================================================================

	// The value shall be the logarithm to the base 2 of the mean announceInterval; see 7.7.2.2.
	LogAnnouncedInterval                       int8

	// The value of portDS.announceReceiptTimeout shall be an integral multiple of announceInterval; see 7.7.3.1.
	// NOTE: The announceInterval is equal to the value of 2^portDS.LogAnnouncedInterval
	AnnounceReceiptTimeout                     uint8

	// The value of portDS.logSyncInterval shall be the logarithm to the base 2 of the mean SyncInterval for multicast
	// messages; see 7.7.2.3.
	// NOTE: The rates for unicast transmissions are negotiated separately on a per port basis and are not constrained
	// by this subclause.
	LogSyncInterval                            int8

	// TODO: Enumeration8
	// The value of portDS.delayMechanism shall indicate the propagation delay measuring option used by the port in
	// computing <meanPathDelay>. The value shall be taken from the enumeration in Table 9. The initialization value is
	// implementation-specific unless otherwise stated in a PTP profile.
	//	DelayMechansim		Value (hex)		Specification
	//	E2E					01				The port is configured to use the delay request-response mechanism.
	//	P2P					02				The port is configured to use the peer delay mechanism.
	//	DISABLED			FE				The port does not implement the delay mechanism; see NOTE.
	//	NOTE: This value shall not be set by a clock except when the applicable PTP profile specifies that the clock
	// 	syntonize only and that neither path delay mechanism is to be used.
	// NOTE: Subclause 9.1 permits reconfiguration. Autoconfiguration is allowed but is out of scope.
	DelayMechanism                             uint8

	// The value of portDS.logMinPdelayReqInterval shall be the logarithm to the base 2 of the minPdelayReqInterval;
	// see 7.7.2.5.
	LogMinPdelayReqInterval                    int8

	// The value of portDS.versionNumber shall indicate the PTP version in use on the port.
	// Version 2
	VersionNumber                            uint8
}



