package utils

import (
	"os"

	"golang.org/x/sys/unix"
)

func GetFreeSpace() (uint64, error) {
	var stat unix.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		return 0, err
	}

	err = unix.Statfs(wd, &stat)
	if err != nil {
		return 0, err
	}

	return stat.Bavail * uint64(stat.Bsize), nil
}
