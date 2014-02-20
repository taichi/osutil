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
package osutil_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/taichi/osutil"
	"io/ioutil"
	"os"
	"path/filepath"
)

var _ = Describe("Osutil", func() {
	Context("Contains", func() {
		It("should work truthy", func() {
			path := "hoge/moge"
			rel := "hoge/moge/../../hoge/moge/piro"
			Expect(Contains(path, rel)).To(BeTrue())
		})
		It("should work falsy", func() {
			path := "hoge/moge"
			rel := "hoge/moge/../../piro"
			Expect(Contains(path, rel)).To(BeFalse())
		})
	})
	Context("IsExist", func() {
		var tmpdir string
		BeforeEach(func() {
			tmpdir, _ = ioutil.TempDir("", "IsExist")
		})
		AfterEach(func() {
			Expect(os.RemoveAll(tmpdir)).To(BeNil())
		})
		It("should work truthy", func() {
			p := filepath.Join(tmpdir, "hoge.txt")
			err := ioutil.WriteFile(p, []byte("aaaaccc"), 0644)
			Expect(err).To(BeNil())
			Expect(IsExist(p)).To(BeTrue())
		})
		It("should work falsy", func() {
			p := filepath.Join(tmpdir, "moge.txt")
			Expect(IsExist(p)).To(BeFalse())
		})
	})
	Context("ForceRemoveAll", func() {
		It("should work normally", func() {
			dir, err := ioutil.TempDir("", "zzz")
			Expect(err).To(BeNil())

			txt := dir + "/moge.txt"
			we := ioutil.WriteFile(txt, []byte("aaaa"), 0644)
			Expect(we).To(BeNil())

			Expect(ForceRemoveAll(dir)).To(BeNil())
			Expect(IsNotExist(dir)).To(BeTrue())
		})
	})
})
