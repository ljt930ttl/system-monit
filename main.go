package main

import "system-monit/info"

func main() {
	table := info.CreatTable()
	info.UITicker(table)
	// info.FmtMemoryInfo()
	// info.FmtDiskInfo()
	// info.FmtNetInfo()
	// unit(0, 1536)

}
