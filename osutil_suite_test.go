package osutil_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/taichi/gotive/ginkgo"

	"testing"
)

func TestOsutil(t *testing.T) {
	RegisterFailHandler(Fail)
	Configure()
	RunSpecs(t, "Osutil Suite")
}
