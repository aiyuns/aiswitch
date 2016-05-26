基于开源项目[swcollector](https://github.com/gaochao1/swcollector)

## 简介
采集的metric列表：
在swcollector的基础上增加了
* hwL2VlanStatInTotalPkts(VLAN下接收报文包数统计)
* hwL2VlanStatOutTotalPkts(VLAN下接收报文字节数统计)
* hwL2VlanStatInTotalBytes(VLAN下发出报文包数统计)
* hwL2VlanStatOutTotalBytes(VLAN下发出报文字节数统计)

## 设备支持
hwL2VlanStatInTotalPkts、hwL2VlanStatOutTotalPkts、hwL2VlanStatInTotalBytes、hwL2VlanStatOutTotalBytes针对华为交换机CE6810(Version 8.80)之用用新标准的oid才能对Vlan进行查询，只在此型号的交换机测试通过，其他版本和型号的交换机请自行测试。

## 源码安装
	mkdir $GOPATH/src/aiyun.com.cn/aiswitch
  cd $GOPATH/src/aiyun.com.cn/aiswitch
  git clone <aiswitch URI>
	// 依赖$GOPATH/src/aiyun.com.cn/aiswitch/sw
	cd $GOPATH/src/aiyun.com.cn/aiswitch/swcollector
	go get ./...
	./control build
	./control pack
	// 最后一步会pack出一个tar.gz的安装包，拿着这个包去部署服务即可。
	// 修改 cfg.json配置文件
	// 启动aiswitch服务
	./control start

	升级时，确保先更新sw
	cd $GOPATH/src/github.com/gaochao1/sw
	git pull

## 部署说明
swcollector需要部署到有交换机SNMP访问权限的服务器上。

使用Go原生的ICMP协议进行Ping探测，swcollector需要root权限运行。

部分交换机使用Go原生SNMP协议会超时。暂时解决方法是SNMP接口流量查询前先判断设备型号，对部分此类设备，调用snmpwalk命令进行数据收集。(一些华为设备和思科的IOS XR)
因此最好在监控探针服务器上也装个snmpwalk命令


## 配置说明
配置文件请参照cfg.example.json，修改该文件名为cfg.json，将该文件里的IP换成实际使用的IP。

switch配置项说明：

	"switch":{
	   "enabled": true,
		"ipRange":[						#交换机IP地址段，对该网段有效IP，先发Ping包探测，对存活IP发SNMP请求
           "192.168.1.0/24",
           "192.168.56.102/32",
           "172.16.114.233"
 		],
		"pingTimeout":300, 			   #Ping超时时间，单位毫秒
		"pingRetry":4,				   #Ping探测重试次数
		"community":"public",			#SNMP认证字符串
		"snmpTimeout":2000,				#SNMP超时时间，单位毫秒
		"snmpRetry":5,					#SNMP重试次数
		"ignoreIface": ["Nu","NU","Vlan","Vl","LoopBack"],    #忽略的接口，如Nu匹配ifName为*Nu*的接口，如果要采集hwVlanStatTotal，请不要忽略Vlan。
		"ignorePkt": true,            #不采集IfHCInUcastPkts、IfHCOutUcastPkts、hwL2VlanStatInTotalPkts、hwL2VlanStatOutTotalPkts
		"ignoreBroadcastPkt": true,   #不采集IfHCInBroadcastPkts和IfHCOutBroadcastPkts
		"ignoreMulticastPkt": true,   #不采集IfHCInMulticastPkts和IfHCOutMulticastPkts
		"ignoreVlanStatTotal": true,  #不采集hwL2VlanStatInTotalPkts、hwL2VlanStatOutTotalPkts、hwL2VlanStatInTotalBytes、hwL2VlanStatOutTotalBytes
		"ignoreOperStatus": true,     #不采集IfOperStatus
		"displayByBit": true,		  #true时，上报的流量单位为bit，为false则单位为byte。
		"fastPingMode": false,	      #是否开启 fastPing 模式，开启 Ping 的效率更高，并能解决高并发时，会有小概率 ping 通宕机的交换机地址的情况。但 fastPing 可能被防火墙过滤。
		"limitConcur": 1000           #限制SNMP请求并发数
    }
