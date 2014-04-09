package main 

import (
	"fmt"
	"alex/ptp1588boundaryclock/datatypes"
	"alex/ptp1588boundaryclock/datasets"
)

const (
		apple = 0xA
		mango = 0xB
		banana = 0xC
	)

func main() {
	fmt.Println("=======================================")
	fmt.Println("Clock Quality")
	clockQuality := new(datatypes.ClockQuality)
	fmt.Println(clockQuality)
	fmt.Println("Clock Identity")
	clockIdentity := new(datatypes.ClockIdentity)
	fmt.Println(clockIdentity)
	
	fmt.Println("=======================================")
	fmt.Println("TLV")
	tlv := new(datatypes.TLV)
	fmt.Println(tlv)
	fmt.Println("PTP Text")
	ptptext := new(datatypes.PTPText)
	fmt.Println(ptptext)
	fmt.Println("Faul Record")
	faultrecord := new(datatypes.FaultRecord)
	fmt.Println(faultrecord)
	
	fmt.Println("=======================================")
	fmt.Println("PortAdress")
	portadress := new(datatypes.PortAdress)
	fmt.Println(portadress)
	fmt.Println("PortIdentity")
	portidentity := new(datatypes.PortIdentity)
	fmt.Println(portidentity)
	
	fmt.Println("=======================================")
	fmt.Println("TimeInterval")
	timeinterval := new(datatypes.TimeInterval)
	fmt.Println(timeinterval)
	fmt.Println("TimeStamp")
	timestamp := new(datatypes.TimeStamp)
	fmt.Println(timestamp)
	
	fmt.Println("=======================================")
	fmt.Println("DefaultDS")
	defaultds := new(datasets.DefaultDS)
	fmt.Println(defaultds)
	
	fmt.Println("=======================================")
	fmt.Println("CurrentDS")
	currentds := new(datasets.CurrentDS)
	fmt.Println(currentds)
	
	fmt.Println("=======================================")
	fmt.Println("ParentDS")
	parentds := new(datasets.ParentDS)
	fmt.Println(parentds)
	
	fmt.Println("=======================================")
	fmt.Println("TimePropertiesDS")
	timepropertiesds := new(datasets.TimePropertiesDS)
	fmt.Println(timepropertiesds)
	
	fmt.Println("=======================================")
	fmt.Println("PortDS")
	portds := new(datasets.PortDS)
	fmt.Println(portds)
	fmt.Println("Hextest: ", 0xFF)
	fmt.Println("Enum: ", apple)
}

