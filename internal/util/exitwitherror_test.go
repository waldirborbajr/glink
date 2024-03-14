package util_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var testT *testing.T

func TestExitWithError(t *testing.T) {
	RegisterFailHandler(Fail)
	testT = t
	RunSpecs(t, "ExitWithError Suite")
}

var _ = Describe("Checking ExitWithError", Label("exiterror"), func() {
	BeforeEach(func() {})

	AfterEach(func() {})

	When("a error is provided", func() {
		It("should exit with error", func() {
		})

		// util.ExitWithError("test message", fmt.Errorf("test error"))
	})
})
