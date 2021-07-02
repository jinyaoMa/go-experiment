package workspace

import (
	"errors"
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

const (
	DISK_ERROR_DRIVE_POINTER = "fail to find drive"
)

var (
	cmd *windows.Proc
)

type Disk struct {
	drive     *uint16
	All       uint64 `json:"all"`
	Used      uint64 `json:"used"`
	Available uint64 `json:"available"`
	Free      uint64 `json:"free"`
}

func InitDisk(drive string) (Disk, error) {
	var disk Disk
	p, err := windows.UTF16PtrFromString(drive)
	if err != nil {
		return disk, errors.New(DISK_ERROR_DRIVE_POINTER)
	}
	if cmd == nil {
		h := windows.MustLoadDLL("kernel32.dll")
		cmd = h.MustFindProc("GetDiskFreeSpaceExW")
	}
	_, _, _ = cmd.Call(uintptr(unsafe.Pointer(p)),
		uintptr(unsafe.Pointer(&disk.Free)),
		uintptr(unsafe.Pointer(&disk.All)),
		uintptr(unsafe.Pointer(&disk.Available)))
	disk.drive = p
	disk.Used = disk.All - disk.Free
	return disk, nil
}

func (disk *Disk) Update() {
	_, _, _ = cmd.Call(uintptr(unsafe.Pointer(disk.drive)),
		uintptr(unsafe.Pointer(&disk.Free)),
		uintptr(unsafe.Pointer(&disk.All)),
		uintptr(unsafe.Pointer(&disk.Available)))
	disk.Used = disk.All - disk.Free
}
