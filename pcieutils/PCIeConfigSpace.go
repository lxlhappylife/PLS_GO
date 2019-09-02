package pcieutils

import (
	log "PLS_GO/log"
	"fmt"
)

//NotSupport8 : Not Support 8 Bits
const NotSupport8 uint8 = 0xFF

//NotSupport16 : Not Support 16 Bits
const NotSupport16 uint16 = 0xFFFF

//NotSupport32 : Not Support 32 Bits
const NotSupport32 uint32 = 0xFFFFFFFF

// GetVID : Get VID
func GetVID(Bus uint8, Device uint8, Function uint8) uint16 {
	return ConfigReadu16(Bus, Device, Function, 0x0)
}

// GetDID : Get DID
func GetDID(Bus uint8, Device uint8, Function uint8) uint16 {
	return ConfigReadu16(Bus, Device, Function, 0x2)
}

// GetClassCode : Get Class Code
func GetClassCode(Bus uint8, Device uint8, Function uint8) uint16 {
	ClassCode := ConfigReadu16(Bus, Device, Function, 0x8) >> 4
	return ClassCode
}

/*
GetMaxLinkSpeed : Return Max Link speed with uint8 and string
*/
func GetMaxLinkSpeed(Bus uint8, Device uint8, Function uint8) (uint8, string) {
	var MaxLinkSpeedStr string
	var linkspeedarray = [7]string{"UNKNOwN", "GEN1", "GEN2", "GEN3", "GEN4", "GEN5", "GEN6"} //Thinking Big and doing small
	pciecapheader := GetPCIeCapHeaderAddress(Bus, Device, Function, 0x10)
	MaxLinkSpeed := 0xF & ConfigReadu32(Bus, Device, Function, uint16(pciecapheader+0xC))
	switch MaxLinkSpeed {
	case 0x1:
		MaxLinkSpeedStr = linkspeedarray[1]
	case 0x2:
		MaxLinkSpeedStr = linkspeedarray[2]
	case 0x3:
		MaxLinkSpeedStr = linkspeedarray[3]
	case 0x4:
		MaxLinkSpeedStr = linkspeedarray[4]
	case 0x5:
		MaxLinkSpeedStr = linkspeedarray[5]
	case 0x6:
		MaxLinkSpeedStr = linkspeedarray[6]
	default:
		MaxLinkSpeedStr = linkspeedarray[0]
	}
	return uint8(MaxLinkSpeed), MaxLinkSpeedStr
}

/*
GetMaxLinkWidth : Return Max Link width with uint8 and string
*/
func GetMaxLinkWidth(Bus uint8, Device uint8, Function uint8) (uint8, string) {
	var MaxLinkWidthStr string
	var linkwidtharray = [7]string{"UNKNOwN", "x1", "x2", "x4", "x8", "x16", "x32"} //Thinking Big and doing small
	pciecapheader := GetPCIeCapHeaderAddress(Bus, Device, Function, 0x10)
	MaxLinkWidth := 0x3F & (ConfigReadu32(Bus, Device, Function, uint16(pciecapheader+0xC)) >> 4)
	switch MaxLinkWidth {
	case 0x1:
		MaxLinkWidthStr = linkwidtharray[1]
	case 0x2:
		MaxLinkWidthStr = linkwidtharray[2]
	case 0x4:
		MaxLinkWidthStr = linkwidtharray[3]
	case 0x8:
		MaxLinkWidthStr = linkwidtharray[4]
	case 0x10:
		MaxLinkWidthStr = linkwidtharray[5]
	case 0x20:
		MaxLinkWidthStr = linkwidtharray[6]
	default:
		MaxLinkWidthStr = linkwidtharray[6]
	}
	return uint8(MaxLinkWidth), MaxLinkWidthStr
}

/*
GetLinkSpeed : Return Link speed with uint8 and string
*/
func GetLinkSpeed(Bus uint8, Device uint8, Function uint8) (uint8, string) {
	var LinkSpeedStr string
	var linkspeedarray = [7]string{"UNKNOwN", "GEN1", "GEN2", "GEN3", "GEN4", "GEN5", "GEN6"} //Thinking Big and doing small
	pciecapheader := GetPCIeCapHeaderAddress(Bus, Device, Function, 0x10)
	LinkSpeed := 0xF & ConfigReadu16(Bus, Device, Function, uint16(pciecapheader+0x12))
	switch LinkSpeed {
	case 0x1:
		LinkSpeedStr = linkspeedarray[1]
	case 0x2:
		LinkSpeedStr = linkspeedarray[2]
	case 0x3:
		LinkSpeedStr = linkspeedarray[3]
	case 0x4:
		LinkSpeedStr = linkspeedarray[4]
	case 0x5:
		LinkSpeedStr = linkspeedarray[5]
	case 0x6:
		LinkSpeedStr = linkspeedarray[6]
	default:
		LinkSpeedStr = linkspeedarray[0]
	}
	return uint8(LinkSpeed), LinkSpeedStr
}

/*
GetLinkWidth : Return Link width with uint8 and string
*/
func GetLinkWidth(Bus uint8, Device uint8, Function uint8) (uint8, string) {
	var LinkWidthStr string
	var linkwidtharray = [7]string{"UNKNOwN", "x1", "x2", "x4", "x8", "x16", "x32"} //Thinking Big and doing small
	pciecapheader := GetPCIeCapHeaderAddress(Bus, Device, Function, 0x10)
	LinkWidth := 0x3F & (ConfigReadu16(Bus, Device, Function, uint16(pciecapheader+0x12)) >> 4)
	switch LinkWidth {
	case 0x1:
		LinkWidthStr = linkwidtharray[1]
	case 0x2:
		LinkWidthStr = linkwidtharray[2]
	case 0x4:
		LinkWidthStr = linkwidtharray[3]
	case 0x8:
		LinkWidthStr = linkwidtharray[4]
	case 0x10:
		LinkWidthStr = linkwidtharray[5]
	case 0x20:
		LinkWidthStr = linkwidtharray[6]
	default:
		LinkWidthStr = linkwidtharray[6]
	}
	return uint8(LinkWidth), LinkWidthStr
}

// GetNumOfRetimer :
func GetNumOfRetimer(Bus uint8, Device uint8, Function uint8) int {
	pciecapheader := GetPCIeCapHeaderAddress(Bus, Device, Function, 0x10)
	var LinkCapabilities2Reg = ConfigReadu32(Bus,
		Device, Function, uint16(pciecapheader+0x2C))
	var LinkStatus2Reg = ConfigReadu16(Bus,
		Device, Function, uint16(pciecapheader+0x32))
	var RetimerPresenceDetectSupported = (LinkCapabilities2Reg >> 23) & 0x1
	var TwoRetimersPresenceDetectSupported = (LinkCapabilities2Reg >> 24) & 0x1
	maxlinkspeed, _ := GetMaxLinkSpeed(Bus, Device, Function)
	if maxlinkspeed < 4 {
		log.Error(fmt.Sprintf("Gen%x does not support retimer.", maxlinkspeed))
		return 0
	}
	if TwoRetimersPresenceDetectSupported == 0x1 {
		TwoRetimersPresenceDetected := (LinkStatus2Reg >> 7) & 0x1
		if TwoRetimersPresenceDetected == 0x1 {
			return 2
		}
	}
	if RetimerPresenceDetectSupported == 0x1 {
		RetimerPresenceDetected := (LinkStatus2Reg >> 6) & 0x1
		if RetimerPresenceDetected == 0x1 {
			return 1
		}
	}
	return 0
}
