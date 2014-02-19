package osutil_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/taichi/gotive/osutil"
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
	Context("HandleWritable", func() {
		var tmpfile *os.File
		BeforeEach(func() {
			var err error
			tmpfile, err = ioutil.TempFile("", "www")
			Expect(err).To(BeNil())
			_, err = tmpfile.WriteString("aaaa")
			Expect(err).To(BeNil())
		})
		AfterEach(func() {
			Expect(tmpfile.Close()).To(BeNil())
			Expect(os.Remove(tmpfile.Name())).To(BeNil())
		})
		It("works normally", func() {
			ToBe := func(perm string) {
				if info, err := os.Lstat(tmpfile.Name()); err != nil {
					Fail(err.Error())
				} else {
					Expect(info.Mode().String()).To(Equal(perm))
				}
			}
			Expect(RemoveWritable(tmpfile)).To(BeNil())
			ToBe("-r--r--r--")
			Expect(AddWritable(tmpfile)).To(BeNil())
			ToBe("-rw-rw-rw-")
		})
	})
	Context("ForceRemoveAll", func() {
		It("should work normally", func() {
			dir, err := ioutil.TempDir("", "zzz")
			Expect(err).To(BeNil())

			txt := dir + "/moge.txt"
			we := ioutil.WriteFile(txt, []byte("aaaa"), 644)
			Expect(we).To(BeNil())

			file, err2 := os.Open(txt)
			Expect(err2).To(BeNil())
			Expect(RemoveWritable(file)).To(BeNil())
			Expect(file.Close()).To(BeNil())
			Expect(os.RemoveAll(dir)).NotTo(BeNil())
			Expect(ForceRemoveAll(dir)).To(BeNil())
		})
	})
})
