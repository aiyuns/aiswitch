swcollector的工具库

snmpwalk -v 2c -c HZidcpal 61.164.47.50 1.3.6.1.4.1.2011.5.25.42.3.1.3.4.1

vlan Name
snmpwalk -v 2c -c HZidcpal 61.164.47.50 1.3.6.1.4.1.2011.5.25.42.3.1.1.1.1.2
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.1 = STRING: "VLAN 0001"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.2 = STRING: "VLAN 0002"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.3 = STRING: "DT:HZ-FY-IDC-NE80E-GuiHua"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.4 = STRING: "VLAN 0004"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.5 = STRING: "VLAN 0005"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.88 = STRING: "VLAN 0088"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.100 = STRING: "TO-HZIDC-FYJQ-9306-ETH-2"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.101 = STRING: "TO-HZIDC-FYJQ-9306-ETH-beiyong"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.102 = STRING: "TO-HZIDC-FYJQ-9306-ETH-3"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.200 = STRING: "VLAN 0200"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.300 = STRING: "to bili-gw"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.630 = STRING: "mb-neiwang-hz-to-sh"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.1000 = STRING: "to yunshao-cnc-out"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.1035 = STRING: "to yunshao-cnc-out"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.1602 = STRING: "HZSW-TO-HZFY"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.1800 = STRING: "TO CN-HZ-SW-OTN-TO-HZFY"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.2501 = STRING: "VLAN 2501"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.2502 = STRING: "VLAN 2502"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.2700 = STRING: "To-C06-bili-HZCUC-20160328"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.3000 = STRING: "TO HZISP NET"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.3001 = STRING: "VLAN 3001"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.3002 = STRING: "VLAN 3002"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.3003 = STRING: "mb-neiwang"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.4000 = STRING: "VLAN 4000"
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.1.1.1.2.4001 = STRING: "VLAN 4001"


SNMPv2-SMI::enterprises.2011.5.25.42.3.1.3.4.1.2.5 = Counter64: 16048938284836
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.3.4.1.3.5 = Counter64: 23777163897178989
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.3.4.1.4.5 = Counter64: 8415779221519
SNMPv2-SMI::enterprises.2011.5.25.42.3.1.3.4.1.5.5 = Counter64: 631340710002442

GAUGE类型
GAUGE翻译过来是计量器的意思。可以理解为最终图表上显示的数据就是采集来的第一手数据。比如气温，停车场中空车位的数目，这种数据随时间变化，并且可上可下，没有固定的规律。
COUNTER类型
COUNTER是计数器的意思。这种类型一般用于记录连续增长的数据，例如某张网卡上流出的数据量。COUNTER数据源会假设计数器的值永远不会减小，除非发生溢出。更新操作会考虑到发生溢出的可能性。计数器会被以每秒的速率存储。当计数器溢出时，RRDtool会检查溢出发生在32bit或是64bit的边界，并对数据加上一个适当的值。

Gauges是一个最简单的计量，一般用来统计瞬时状态的数据信息，比如系统中处于pending状态的job。

http://blog.csdn.net/smallnest/article/details/38491507

iso(1).org(3).dod(6).internet(1).private(4).enterprises(1).huawei(2011).huaweiMgmt(5).hwDatacomm(25).hwLldpMIB(134)

ifInMulticastPkts [2]
 Counter32
 read-only
 接收的组播报文个数。

对MAC层协议来说，组播地址包含了组地址和功能地址。
 实现与MIB文件定义一致。

ifInBroadcastPkts [3]
 Counter32
 read-only
 接收的广播报文个数。
 实现与MIB文件定义一致。

ifOutMulticastPkts [4]
 Counter32
 read-only
 发送的组播报文总数，包括丢弃的报文和没有发送的报文。

对MAC层协议来说，组播地址包含了组地址和功能地址。
 实现与MIB文件定义一致。

ifOutBroadcastPkts [5]
 Counter32
 read-only
 发送的广播报文总数，包括被丢弃的报文或没有发送的报文。
 实现与MIB文件定义一致。

ifHCInOctets [6]
 Counter64
 read-only
 接口上接收到的字节总数，包括成帧的字符。该节点有64bit，是ifInOctets的扩充。
 实现与MIB文件定义一致。

ifHCInUcastPkts [7]
 Counter64
 read-only
 接口上接收到的单播报文个数。该节点是ifInUcastPkts的扩充，有64bit。
 实现与MIB文件定义一致。

ifHCInMulticastPkts [8]
 Counter64
 read-only
 接收的组播报文个数。对于MAC层协议来说，组播地址包括组地址和功能地址。该节点是ifInMulticastPkts的扩充，有64bit。
 实现与MIB文件定义一致。

ifHCInBroadcastPkts [9]
 Counter64
 read-only
 接收的广播报文个数。该节点是ifInBroadcastPkts的扩充，有64bit。
 实现与MIB文件定义一致。

ifHCOutOctets [10]
 Counter64
 read-only
 接口发送的字节总数，包括成帧字符。该节点是ifOutOctets的扩充，有64bit。
 实现与MIB文件定义一致。

ifHCOutUcastPkts [11]
 Counter64
 read-only
 发送的单播报文总数，包括被丢弃的报文或没有送出的报文。该节点是ifOutUcastPkts的扩充，有64bit。
 实现与MIB文件定义一致。

ifHCOutMulticastPkts [12]
 Counter64
 read-only
 发送的组播报文总数，包括被丢弃的报文或没有送出的报文。对于MAC层协议，组播地址包括组地址和功能地址。该节点是ifOutMulticastPkts的扩充，有64bit。
 实现与MIB文件定义一致。

ifHCOutBroadcastPkts [13]
 Counter64
 read-only
 发送的广播报文总数，包括被丢弃的报文或没有送出的报文。该节点是ifOutBroadcastPkts的扩充，有64bit。
 实现与MIB文件定义一致。
