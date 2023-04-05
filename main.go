package main

import (
	"system-monit/info"
)

func main() {
	// info.GetConnents()
	table := info.CreatTableWriter()
	table.UITicker()
	// info.FmtMemoryInfo()
	// info.FmtDiskInfo()
	// info.FmtNetInfo()
	// unit(0, 1536)

}
