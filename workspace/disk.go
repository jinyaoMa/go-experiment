package workspace

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	// B 1bytes
	B = 1
	// KB 1024bytes
	KB = 1024 * B
	// MB 1024 * 1024bytes
	MB = 1024 * KB
	// GB 1024 * 1024 * 1024bytes
	GB = 1024 * MB
)

type Disk struct {
	All       uint64 `json:"all"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
	Free      uint64 `json:"free"`
}

func InitDisk(drive string) Disk {
	var disk Disk
	h := windows.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	_, _, _ = c.Call(uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(drive))),
		uintptr(unsafe.Pointer(&disk.Free)),
		uintptr(unsafe.Pointer(&disk.All)),
		uintptr(unsafe.Pointer(&disk.Available)))
	disk.Used = disk.All - disk.Free
	return disk
}
