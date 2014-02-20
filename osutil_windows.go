/* Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package osutil

import (
	"os"
	"path/filepath"
	"syscall"
)

func forceRemoveAll(root string) error {
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
			// see. http://msdn.microsoft.com/en-us/library/windows/desktop/aa365535(v=vs.85).aspx
			return syscall.Chmod(f.Name(), syscall.S_IWRITE)
		}
	}
	return nil
}
