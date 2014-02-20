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
	return forceRemoveAll(root)
}
