package info

import "github.com/shirou/gopsutil/net"

type NetInfo struct {
	Name        string `json:"name"`
	BytesSent   uint64 `json:"bytesSent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytesRecv"`   // number of bytes received
	PacketsSent uint64 `json:"packetsSent"` // number of packets sent
	PacketsRecv uint64 `json:"packetsRecv"` // number of packets received
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
