package info

import (
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func CreatTable() table.Writer {
	newTable := table.NewWriter()
	// newTable.SetAutoIndex(true)
	newTable.Style().Options.SeparateHeader = true
	newTable.Style().Options.SeparateRows = true
	newTable.Style().Options.SeparateFooter = true
	newTable.Style().Title.Align = text.AlignCenter
	return newTable
}
func FmtPercent(t table.Writer, name string) {

	//字体颜色
	warnColor := text.Colors{text.BgRed}
	warnTransFormer := text.Transformer(func(val interface{}) string {
		if val.(float64) > 80 {
			return warnColor.Sprintf("%.2f%%", val)
		}
		return fmt.Sprintf("%.2f%%", val)
	})
	t.Style().Format = table.FormatOptions{
		Header: text.FormatTitle,
	}
	t.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:        name,
			Align:       text.AlignCenter,
			AlignHeader: text.AlignCenter,
			AlignFooter: text.AlignCenter,
			Transformer: warnTransFormer,
		},
	})

}

func resetTable(t table.Writer) {
	t.ResetRows()
	t.ResetHeaders()
}

func updateHost(t table.Writer) {
	hInfo := GetHostInfo()
	resetTable(t)
	t.SetTitle("系统信息")
	t.AppendRows([]table.Row{
		{"系统名称", hInfo.Hostname},
		{"系统类别", hInfo.OS},
		{"系统类型", hInfo.KernelArch},
		{"发行版本", fmt.Sprintf("%s-%s", hInfo.Platform, hInfo.PlatformVersion)},
		{"内核版本", hInfo.KernelVersion},
		{"启动时间", time.Unix(int64(hInfo.BootTime), 0).Format(time.DateTime)},
		{"运行时间", fmtSinceTime(time.Unix(int64(hInfo.Uptime), 0))},
	})
}
func updateCPU(t table.Writer) {
	cpuInfo := GetCpuInfo()
	cpuCurrentInfo := GetCpuCurrentInfo(cpuInfo)
	loadInfo := GetLoadCurrentInfo(cpuInfo)
	resetTable(t)
	t.SetTitle("cpu信息")
	t.AppendHeader(table.Row{fmt.Sprintf("Cores:%d, Logical Cores:%d", cpuInfo.CPUCores, cpuInfo.CPULogicalCores), cpuInfo.ModelName})

	cTable := CreatTable()
	// cTable.SetStyle(table.StyleColoredBlackOnCyanWhite)
	cTable.Style().Title.Align = text.AlignCenter
	FmtPercent(cTable, "Percent")
	cTable.SetTitle("Userd")
	cTable.AppendHeader(table.Row{"coreID", "Percent"})
	for idx, percent := range cpuCurrentInfo.CouesUsedPercent {
		cTable.AppendRow(table.Row{idx, percent})
	}
	cTable.AppendRow(table.Row{"Avg", cpuCurrentInfo.CPUAvgPercent})
	cTable.AppendRow(table.Row{"Total", cpuCurrentInfo.CPUTotalUsePercent})

	lTable := CreatTable()
	FmtPercent(lTable, "Percent")
	// lTable.SetStyle(table.StyleColoredDark)
	lTable.Style().Title.Align = text.AlignCenter
	lTable.SetTitle("Load")
	lTable.AppendHeader(table.Row{"Load", "Percent"})
	lTable.AppendRow(table.Row{"Load1", loadInfo.Load1 * 100})
	lTable.AppendRow(table.Row{"Load5", loadInfo.Load5 * 100})
	lTable.AppendRow(table.Row{"Load15", loadInfo.Load15 * 100})
	lTable.AppendRow(table.Row{"Load1-UsagePercent", loadInfo.UsagePercent})

	t.AppendRow(table.Row{cTable.Render(), lTable.Render()})
}

func updateMemory(t table.Writer) {

	MemInfoLast := GetMemInfo()
	SwapMemInfoLast := GetSwapInfo()
	t.SetTitle("内存信息")
	resetTable(t)
	t.AppendHeader(table.Row{"type", "total", "used", "userPercent", "free", "available"})
	t.AppendRow(
		table.Row{
			MemInfoLast.Type,
			convertUnit(B, float64(MemInfoLast.Total)),
			convertUnit(B, float64(MemInfoLast.Used)),
			MemInfoLast.UsedPercent,
			convertUnit(B, float64(MemInfoLast.Free)),
			convertUnit(B, float64(MemInfoLast.Available)),
		})
	t.AppendRow(
		table.Row{
			SwapMemInfoLast.Type,
			convertUnit(B, float64(SwapMemInfoLast.Total)),
			convertUnit(B, float64(SwapMemInfoLast.Used)),
			SwapMemInfoLast.UsedPercent,
			convertUnit(B, float64(SwapMemInfoLast.Free)),
		})

	FmtPercent(t, "userPercent")

}

func updateDisk(t table.Writer) {
	rows, err := GetDiskInfo()
	if err != nil {
		fmt.Printf("获取磁盘信息失败 , err:%v\n", err)
	}
	resetTable(t)
	t.SetTitle("硬盘信息")
	// fmt.Printf("%v", rows)
	dTable := CreatTable()
	dTable.SetTitle("device")
	dTable.AppendHeader(table.Row{"Device", "fstype", "path", "total", "used", "free", "userPercent"})
	iTable := CreatTable()
	iTable.SetTitle("Inodes")
	iTable.AppendHeader(table.Row{"Device", "total", "used", "free", "userPercent"})

	for _, row := range rows {
		dTable.AppendRow(
			table.Row{
				row.Device,
				row.Fstype,
				row.Path,
				convertUnit(B, float64(row.Total)),
				convertUnit(B, float64(row.Used)),
				convertUnit(B, float64(row.Free)),
				row.UsedPercent,
			})

		iTable.AppendRow(
			table.Row{
				row.Device,
				convertUnit(B, float64(row.InodesTotal)),
				convertUnit(B, float64(row.InodesUsed)),
				convertUnit(B, float64(row.InodesFree)),
				row.InodesUsedPercent,
			})
	}
	FmtPercent(dTable, "userPercent")
	FmtPercent(iTable, "userPercent")

	t.AppendRow(table.Row{dTable.Render(), iTable.Render()})
}

func updateNet(t table.Writer) {
	nets, err := GetNetIO()
	if err != nil {
		fmt.Printf("获取网卡信息失败 , err:%v\n", err)
	}
	t.SetTitle("网络信息")
	resetTable(t)
	t.AppendHeader(table.Row{"Name", "BytesSent", "BytesRecv", "PacketsSent", "PacketsRecv"})
	for _, row := range nets {
		t.AppendRow(table.Row{row.Name, row.BytesSent, row.BytesRecv, row.PacketsSent, row.PacketsRecv})
	}

}

func render(t table.Writer) {
	// fmt.Fprintln(u.Write, t.Render())
	clear()
	fmt.Print(t.Render())
}

func UITicker(t table.Writer) {
	var a string
	ticker := time.NewTicker(time.Second * 2) // 创建一个定时器对象
	fmt.Print("start...\ncpu信息输入c，内存信息输入m，硬盘信息输入d，网络信息输入n")
	go func() {
		for {
			select {
			case <-ticker.C: // 每隔一秒，会执行一次
				switch a {
				case "c":
					updateCPU(t)
					render(t)
				case "d":
					updateDisk(t)
					render(t)
				case "h":
					updateHost(t)
					render(t)
				case "n":
					updateNet(t)
					render(t)
				case "m":
					updateMemory(t)
					render(t)
				case "q":
					ticker.Stop()
				default:
					updateCPU(t)
					render(t)
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
