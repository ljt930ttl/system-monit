package info

import (
	"github.com/shirou/gopsutil/net"
)

type NetInfo struct {
	Name        string `json:"name"`
	BytesSent   uint64 `json:"bytesSent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytesRecv"`   // number of bytes received
	PacketsSent uint64 `json:"packetsSent"` // number of packets sent
	PacketsRecv uint64 `json:"packetsRecv"` // number of packets received
}

type ConnentInfo struct {
	Fd     uint32 `json:"fd"`
	Family uint32 `json:"family"`
	Type   string `json:"type"` //TCP:1, UPD:2
	LIP    string `json:"lip"`
	LPort  uint32 `json:"lport"`
	RIP    string `json:"rip"`
	RPort  uint32 `json:"rport"`
	Status string `json:"status"`
	Pid    int32  `json:"pid"`
}

var sockTypeMap = map[uint32]string{
	1: "TCP",
	2: "UDP",
}

func GetNetIO() ([]NetInfo, error) {
	netInfos := []NetInfo{}
	netTotalInfo := NetInfo{}
	netStat, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	for _, stat := range netStat {
		netTotalInfo.BytesSent += stat.BytesSent
		netTotalInfo.BytesRecv += stat.BytesRecv
		netTotalInfo.PacketsSent += stat.PacketsSent
		netTotalInfo.PacketsRecv += stat.PacketsRecv

		netInfos = append(netInfos, NetInfo{
			Name:        stat.Name,
			BytesSent:   stat.BytesSent,
			BytesRecv:   stat.BytesRecv,
			PacketsSent: stat.PacketsSent,
			PacketsRecv: stat.PacketsRecv,
		})
	}
	netTotalInfo.Name = "ALL"
	netInfos = append(netInfos, netTotalInfo)
	return netInfos, err
}

func GetConnents() []ConnentInfo {
	connInfoALL := make([]ConnentInfo, 0)
	conns, _ := net.Connections("inet4")
	for _, conn := range conns {
		connInfo := ConnentInfo{}
		connInfo.Fd = conn.Fd
		connInfo.Family = conn.Family
		connInfo.Type = sockTypeMap[conn.Type]
		connInfo.LIP = conn.Laddr.IP
		connInfo.LPort = conn.Laddr.Port
		connInfo.RIP = conn.Raddr.IP
		connInfo.RPort = conn.Raddr.Port
		connInfo.Status = conn.Status
		connInfo.Pid = conn.Pid
		connInfoALL = append(connInfoALL, connInfo)
	}
	return connInfoALL
}
