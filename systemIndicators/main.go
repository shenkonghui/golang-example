package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)


func main() {
	fmt.Println(GetCpuPercent())
	fmt.Println(GetMemPercent())

	fmt.Println(GetMemFree())
	fmt.Println(GetDiskFree())
}

func GetCpuPercent() float64 {
	percent, _:= cpu.Percent(time.Second, false)
	return percent[0]
}

func GetMemPercent()float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetMemFree() uint64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.Free
}

func GetDiskFree() uint64{
	parts, _ := disk.Partitions(false)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.Free
}



func GetDiskCapacity() uint64{
	parts, _ := disk.Partitions(false)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.Free
}