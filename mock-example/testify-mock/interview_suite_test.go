package testify

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mock Example Suite")
}
