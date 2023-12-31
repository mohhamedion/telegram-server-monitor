package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	"telegram_server_monitor/reporter"
	"telegram_server_monitor/telegrambot"

	"syscall"
)

func checkFreeSpace() (uint64, error) {
	path := "/" // Change this to the directory or drive you want to check
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

	reporterInstence := reporter.GetInstence()

	// later, we add more layers, such as Slack or db
	reporterInstence.AddLayer(&telegrambot.TelegramInformer{})
	var memStats runtime.MemStats

	for {
		runtime.ReadMemStats(&memStats)

		memoryInfo, err := mem.VirtualMemory()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		freeSpace, err := checkFreeSpace()

		percentages, err := cpu.Percent(time.Second, false)

		reporterInstence.Inform(
			fmt.Sprintf("Memory Used : %.2f%%\n", memoryInfo.UsedPercent) +
				fmt.Sprintf("CPU Usage: %.2f%%\n", percentages[0]) +
				fmt.Sprintf("Free space on: %v mb\n", freeSpace))

		time.Sleep(5 * time.Second)

	}
}
