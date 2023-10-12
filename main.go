package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	"telegram_server_monitor/telegrambot"

	"syscall"
)

func checkFreeSpace(path string) (uint64, error) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return 0, err
	}

	// Calculate free space in megabytes
	freeSpaceInBytes := stat.Bfree * uint64(stat.Bsize)
	freeSpaceInMB := float64(freeSpaceInBytes) / 1048576

	return uint64(freeSpaceInMB), nil
}

func main() {
	// Get memory usage information

	path := "/" // Change this to the directory or drive you want to check

	for {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		memoryInfo, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		freeSpace, err := checkFreeSpace(path)

		percentages, err := cpu.Percent(time.Second, false)

		telegrambot.Inform(
			fmt.Sprintf("Memory Used : %.2f%%\n", memoryInfo.UsedPercent) +
				fmt.Sprintf("CPU Usage: %.2f%%\n", percentages[0]) +
				fmt.Sprintf("Free space on %s: %d mb\n", path, freeSpace))

		time.Sleep(5 * time.Second)

	}
}
