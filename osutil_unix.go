// +build darwin dragonfly freebsd linux netbsd openbsd

package osutil

import (
	"os"
)

func addWritable(file *os.File) error {
	return handleWritable(func(v os.FileMode) uint32 { return v | 0222 })
}

func removeWritable(file *os.File) error {
	return handleWritable(func(v os.FileMode) uint32 { return v ^ 0222 })
}

func handleWritable(file *os.File, fn func(perm os.FileMode) uint32) error {
	if info, err := os.Lstat(file.Name()); err != nil {
		return err
	} else {
		return file.Chmod(fn(info.Mode()))
	}
}
