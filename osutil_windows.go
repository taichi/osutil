package osutil

import (
	"os"
	"syscall"
)

// see. http://msdn.microsoft.com/en-us/library/windows/desktop/aa365535(v=vs.85).aspx
func addWritable(file *os.File) error {
	return syscall.Chmod(file.Name(), syscall.S_IWRITE)
}
func removeWritable(file *os.File) error {
	return syscall.Chmod(file.Name(), 01)
}
