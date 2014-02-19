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
