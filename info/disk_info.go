package info

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/disk"
)

type DiskInfo struct {
	Device            string  `json:"device"`
	Path              string  `json:"path"`
	Fstype            string  `json:"fstype"`
	Opts              string  `json:"opts"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"usedPercent"`
	InodesTotal       uint64  `json:"inodesTotal"`
	InodesUsed        uint64  `json:"inodesUsed"`
	InodesFree        uint64  `json:"inodesFree"`
	InodesUsedPercent float64 `json:"inodesUsedPercent"`
}

type DiskIO struct {
	IOReadBytes  uint64
	IOWriteBytes uint64
	IOCount      uint64
	IOTime       uint64
}

// 磁盘信息
func GetDiskInfo() ([]DiskInfo, error) {
	diskInfoToal := DiskInfo{}
	diskInfoRows := []DiskInfo{}

	parts, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	for _, part := range parts {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			fmt.Printf("获取磁盘信息失败 , err:%v\n", err)
			continue
		}
		diskInfo := DiskInfo{}
		diskInfo.Device = part.Device
		diskInfo.Path = part.Mountpoint
		diskInfo.Fstype = part.Fstype
		diskInfo.Total = usage.Total
		diskInfo.Used = usage.Used
		diskInfo.Free = usage.Free
		diskInfo.UsedPercent = usage.UsedPercent
		diskInfo.InodesTotal = usage.InodesTotal
		diskInfo.InodesUsed = usage.InodesUsed
		diskInfo.InodesFree = usage.InodesFree
		diskInfo.InodesUsedPercent = usage.InodesUsedPercent

		diskInfoRows = append(diskInfoRows, diskInfo)
		diskInfoToal.Total += diskInfo.Total
		diskInfoToal.Used += diskInfo.Used
		diskInfoToal.Free += diskInfo.Free
		diskInfoToal.InodesTotal += diskInfo.InodesTotal
		diskInfoToal.InodesUsed += diskInfo.InodesUsed
		diskInfoToal.InodesFree += diskInfo.InodesFree
		// fmt.Printf("该磁盘分区{%v}使用信息:Total:%v\t used:%v\t free:%v\t usedPercent:%0.2f%%\t\n",
		// 	part.Mountpoint, usage.Total/1024/1024/1024, usage.Used/1024/1024/1024,
		// 	usage.Free/1024/1024/1024, float64(usage.Used)/float64(usage.Total)*100)
	}
	diskInfoToal.Device = "All"
	diskInfoToal.UsedPercent = (float64(diskInfoToal.Total) - float64(diskInfoToal.Free)) / float64(diskInfoToal.Total) * 100
	diskInfoToal.InodesUsedPercent = (float64(diskInfoToal.InodesTotal) - float64(diskInfoToal.InodesFree)) / float64(diskInfoToal.InodesTotal) * 100
	diskInfoRows = append(diskInfoRows, diskInfoToal)

	return diskInfoRows, nil
	// fmt.Printf("\n该磁盘使用信息:Total:%v\t used:%v\t free:%v\t\n", diskInfoToal.Total, diskInfoToal.Used, diskInfoToal.Free)
}

func GetDiskIO() *DiskIO {
	DiskIOLast := &DiskIO{}

	ioStat, _ := disk.IOCounters()
	for _, state := range ioStat {
		DiskIOLast.IOReadBytes += state.ReadBytes
		DiskIOLast.IOWriteBytes += state.WriteBytes
		DiskIOLast.IOCount += (state.ReadCount + state.WriteCount)
		DiskIOLast.IOTime += (state.ReadTime + state.WriteTime)
	}
	return DiskIOLast
}

func ContinueIO() {
	var a string
	fmt.Printf("disk io\n ")
	ticker := time.NewTicker(time.Second * 1) // 创建一个定时器对象
	//开启go协程
	go func() {
		for {
			select {
			case <-ticker.C: // 每隔一秒，会执行一次
				fmt.Printf(" .....\n")
				// i++
				if a == "q" {
					ticker.Stop()
				}
			}

		}

	}()

	for {
		fmt.Scanf("%s", &a)
		if a == "q" {
			fmt.Println("Bay Bay ")
			break
		}
	}

}
