package osutil

import (
	"os"
	"path/filepath"
	"strings"
)

func Contains(basepath, relativepath string) bool {
	p, err := filepath.Rel(basepath, relativepath)
	return err == nil && strings.HasPrefix(p, "..") == false
}

func IsExist(path string) bool {
	return IsNotExist(path) == false
}
func IsNotExist(path string) bool {
	_, err := os.Lstat(path)
	return os.IsNotExist(err)
}

func ForceRemoveAll(root string) error {
	// cf.
	// DeleteFile function
	//    http://msdn.microsoft.com/en-us/library/windows/desktop/aa363915(v=vs.85).aspx
	// SHFileOperation function
	//    http://msdn.microsoft.com/en-us/library/bb762164(v=vs.85).aspx
	// IFileOperation interface
	//    http://msdn.microsoft.com/en-us/library/bb775771(v=vs.85).aspx
	if err := filepath.Walk(root, addWritables); err != nil {
		return err
	}
	return os.RemoveAll(root)
}

func addWritables(path string, info os.FileInfo, err error) error {
	if info.Mode()&0200 == 0 {
		if f, err := os.Open(path); err != nil {
			return err
		} else {
			defer f.Close()
			return addWritable(f)
		}
	}
	return nil
}
