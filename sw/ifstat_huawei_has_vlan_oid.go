package sw

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gaochao1/gosnmp"
)

const (
	hwL2VlanStatInTotalPktsOid        = "1.3.6.1.4.1.2011.5.25.42.3.1.3.4.1.2"
	hwL2VlanStatInTotalPktsOidPrefix  = ".1.3.6.1.4.1.2011.5.25.42.3.1.3.4.1.2."
	hwL2VlanStatOutTotalPktsOid       = "1.3.6.1.4.1.2011.5.25.42.3.1.3.4.1.4"
	hwL2VlanStatInTotalBytesOid       = "1.3.6.1.4.1.2011.5.25.42.3.1.3.4.1.3"
	hwL2VlanStatInTotalBytesOidPrefix = ".1.3.6.1.4.1.2011.5.25.42.3.1.3.4.1.3."
	hwL2VlanStatOutTotalBytesOid      = "1.3.6.1.4.1.2011.5.25.42.3.1.3.4.1.5"
)

func ListIfStatsHuaweiHasVlanOid(ip, community string, timeout int, ignoreIface []string, retry int, ignorePkt, ignoreOperStatus, ignoreBroadcastPkt, ignoreMulticastPkt, ignoreVlanStatTotal bool) ([]IfStats, error) {
	var ifStatsList []IfStats

	defer func() {
		if r := recover(); r != nil {
			log.Println(ip+" Recovered in ListIfStats", r)
		}
	}()

	chIfInList := make(chan []gosnmp.SnmpPDU)
	chIfOutList := make(chan []gosnmp.SnmpPDU)

	chIfNameList := make(chan []gosnmp.SnmpPDU)
	chIfStatusList := make(chan []gosnmp.SnmpPDU)

	go ListIfHCInOctets(ip, community, timeout, chIfInList, retry)
	go ListIfHCOutOctets(ip, community, timeout, chIfOutList, retry)
	go ListIfName(ip, community, timeout, chIfNameList, retry)

	ifInList := <-chIfInList
	ifOutList := <-chIfOutList

	ifNameList := <-chIfNameList

	chIfInPktList := make(chan []gosnmp.SnmpPDU)
	chIfOutPktList := make(chan []gosnmp.SnmpPDU)

	var ifInPktList, ifOutPktList []gosnmp.SnmpPDU

	if ignorePkt == false {
		go ListIfHCInUcastPkts(ip, community, timeout, chIfInPktList, retry)
		go ListIfHCOutUcastPkts(ip, community, timeout, chIfOutPktList, retry)
		ifInPktList = <-chIfInPktList
		ifOutPktList = <-chIfOutPktList
	}

	chIfInBroadcastPktList := make(chan []gosnmp.SnmpPDU)
	chIfOutBroadcastPktList := make(chan []gosnmp.SnmpPDU)

	var ifInBroadcastPktList, ifOutBroadcastPktList []gosnmp.SnmpPDU

	if ignoreBroadcastPkt == false {
		go ListIfHCInBroadcastPkts(ip, community, timeout, chIfInBroadcastPktList, retry)
		go ListIfHCOutBroadcastPkts(ip, community, timeout, chIfOutBroadcastPktList, retry)
		ifInBroadcastPktList = <-chIfInBroadcastPktList
		ifOutBroadcastPktList = <-chIfOutBroadcastPktList
	}

	chIfInMulticastPktList := make(chan []gosnmp.SnmpPDU)
	chIfOutMulticastPktList := make(chan []gosnmp.SnmpPDU)

	var ifInMulticastPktList, ifOutMulticastPktList []gosnmp.SnmpPDU

	if ignoreMulticastPkt == false {
		go ListIfHCInMulticastPkts(ip, community, timeout, chIfInMulticastPktList, retry)
		go ListIfHCOutMulticastPkts(ip, community, timeout, chIfOutMulticastPktList, retry)
		ifInMulticastPktList = <-chIfInMulticastPktList
		ifOutMulticastPktList = <-chIfOutMulticastPktList
	}

	var ifStatusList []gosnmp.SnmpPDU
	if ignoreOperStatus == false {
		go ListIfOperStatus(ip, community, timeout, chIfStatusList, retry)
		ifStatusList = <-chIfStatusList
	}

	chHWL2VlanStatInTotalPktsList := make(chan []gosnmp.SnmpPDU)
	chHWL2VlanStatInTotalBytesList := make(chan []gosnmp.SnmpPDU)
	chHWL2VlanStatOutTotalPktsList := make(chan []gosnmp.SnmpPDU)
	chHWL2VlanStatoutTotalBytesList := make(chan []gosnmp.SnmpPDU)

	var (
		hwL2VlanStatInTotalPktsList   []gosnmp.SnmpPDU
		hwL2VlanStatInTotalBytesList  []gosnmp.SnmpPDU
		hwL2VlanStatOutTotalPktsList  []gosnmp.SnmpPDU
		hwL2VlanStatOutTotalBytesList []gosnmp.SnmpPDU
	)
	// ps. 增加对Vlan端口的独立统计
	// VLAN下接收报文包数统计
	// VLAN下接收报文字节数统计
	// VLAN下发出报文包数统计
	// VLAN下发出报文字节数统计
	if !ignoreVlanStatTotal {
		go ListHWL2VlanStatInTotalPkts(ip, community, timeout, chHWL2VlanStatInTotalPktsList, retry)
		go ListHWL2VlanStatInTotalBytes(ip, community, timeout, chHWL2VlanStatInTotalBytesList, retry)
		go ListHWL2VlanStatOutTotalPkts(ip, community, timeout, chHWL2VlanStatOutTotalPktsList, retry)
		go ListHWL2VlanStatOutTotalBytes(ip, community, timeout, chHWL2VlanStatoutTotalBytesList, retry)
		hwL2VlanStatInTotalPktsList = <-chHWL2VlanStatInTotalPktsList
		hwL2VlanStatInTotalBytesList = <-chHWL2VlanStatInTotalBytesList
		hwL2VlanStatOutTotalPktsList = <-chHWL2VlanStatOutTotalPktsList
		hwL2VlanStatOutTotalBytesList = <-chHWL2VlanStatoutTotalBytesList
	}

	if len(ifNameList) > 0 && len(ifInList) > 0 && len(ifOutList) > 0 {
		now := time.Now().Unix()

		for _, ifNamePDU := range ifNameList {

			ifName := ifNamePDU.Value.(string)

			check, checkVlanStatTotal := true, false
			if len(ignoreIface) > 0 {
				for _, ignore := range ignoreIface {
					if strings.Contains(ifName, ignore) {
						check = false
						break
					}
				}
			}

			if strings.Contains(ifName, "Vlan") {
				checkVlanStatTotal = true
			}

			if check {
				var ifStats IfStats

				ifIndexStr := strings.Replace(ifNamePDU.Name, ifNameOidPrefix, "", 1)

				ifStats.IfIndex, _ = strconv.Atoi(ifIndexStr)

				for ti, ifHCInOctetsPDU := range ifInList {
					if strings.Replace(ifHCInOctetsPDU.Name, ifHCInOidPrefix, "", 1) == ifIndexStr {
						ifStats.IfHCInOctets = ifInList[ti].Value.(uint64)
						ifStats.IfHCOutOctets = ifOutList[ti].Value.(uint64)
					}
				}
				if ignorePkt == false {
					for ti, ifHCInPktsPDU := range ifInPktList {
						if strings.Replace(ifHCInPktsPDU.Name, ifHCInPktsOidPrefix, "", 1) == ifIndexStr {
							ifStats.IfHCInUcastPkts = ifInPktList[ti].Value.(uint64)
							ifStats.IfHCOutUcastPkts = ifOutPktList[ti].Value.(uint64)
						}
					}
				}
				if ignoreBroadcastPkt == false {
					for ti, ifHCInBroadcastPktPDU := range ifInBroadcastPktList {
						if strings.Replace(ifHCInBroadcastPktPDU.Name, ifHCInBroadcastPktsOidPrefix, "", 1) == ifIndexStr {
							ifStats.IfHCInBroadcastPkts = ifInBroadcastPktList[ti].Value.(uint64)
							ifStats.IfHCOutBroadcastPkts = ifOutBroadcastPktList[ti].Value.(uint64)
						}
					}
				}
				if ignoreMulticastPkt == false {
					for ti, ifHCInMulticastPktPDU := range ifInMulticastPktList {
						if strings.Replace(ifHCInMulticastPktPDU.Name, ifHCInMulticastPktsOidPrefix, "", 1) == ifIndexStr {
							ifStats.IfHCInMulticastPkts = ifInMulticastPktList[ti].Value.(uint64)
							ifStats.IfHCOutMulticastPkts = ifOutMulticastPktList[ti].Value.(uint64)
						}
					}
				}
				if ignoreOperStatus == false {
					for ti, ifOperStatusPDU := range ifStatusList {
						if strings.Replace(ifOperStatusPDU.Name, ifOperStatusOidPrefix, "", 1) == ifIndexStr {
							ifStats.IfOperStatus = ifStatusList[ti].Value.(int)
						}
					}
				}
				// ps. 把wlan的统计放入返回的对象
				if !ignoreVlanStatTotal && checkVlanStatTotal {
					for ti, hwL2VlanStatTotalBytesPDU := range hwL2VlanStatInTotalBytesList {
						if strings.HasSuffix(ifName, strings.Replace(hwL2VlanStatTotalBytesPDU.Name, hwL2VlanStatInTotalBytesOidPrefix, "", 1)) {
							ifStats.HWL2VlanStatInTotalBytes = hwL2VlanStatInTotalBytesList[ti].Value.(uint64)
							ifStats.HWL2VlanStatOutTotalBytes = hwL2VlanStatOutTotalBytesList[ti].Value.(uint64)
						}
					}
				}
				if !ignoreVlanStatTotal && checkVlanStatTotal {
					for ti, hwL2VlanStatTotalPktsPDU := range hwL2VlanStatInTotalPktsList {
						if strings.HasSuffix(ifName, strings.Replace(hwL2VlanStatTotalPktsPDU.Name, hwL2VlanStatInTotalPktsOidPrefix, "", 1)) {
							ifStats.HWL2VlanStatInTotalPkts = hwL2VlanStatInTotalPktsList[ti].Value.(uint64)
							ifStats.HWL2VlanStatOutTotalPkts = hwL2VlanStatOutTotalPktsList[ti].Value.(uint64)
						}
					}
				}
				ifStats.TS = now
				ifStats.IfName = ifName
				ifStatsList = append(ifStatsList, ifStats)
			}
		}
	}

	return ifStatsList, nil
}

func ListHWL2VlanStatInTotalPkts(ip, community string, timeout int, ch chan []gosnmp.SnmpPDU, retry int) {
	RunSnmpRetry(ip, community, timeout, ch, retry, hwL2VlanStatInTotalPktsOid)
}

func ListHWL2VlanStatInTotalBytes(ip, community string, timeout int, ch chan []gosnmp.SnmpPDU, retry int) {
	RunSnmpRetry(ip, community, timeout, ch, retry, hwL2VlanStatInTotalBytesOid)
}

func ListHWL2VlanStatOutTotalPkts(ip, community string, timeout int, ch chan []gosnmp.SnmpPDU, retry int) {
	RunSnmpRetry(ip, community, timeout, ch, retry, hwL2VlanStatOutTotalPktsOid)
}

func ListHWL2VlanStatOutTotalBytes(ip, community string, timeout int, ch chan []gosnmp.SnmpPDU, retry int) {
	RunSnmpRetry(ip, community, timeout, ch, retry, hwL2VlanStatOutTotalBytesOid)
}
