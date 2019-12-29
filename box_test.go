package refiment

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRTBox_UnpackPositiveTest(t *testing.T) {
	g := NewGomegaWithT(t)

	testValue := 5
	testBox := NewBox(validMapFunc, testValue)
	unBoxed, err := testBox.Unpack()

	g.Expect(unBoxed).To(Equal(testValue))
	g.Expect(err).ShouldNot(HaveOccurred())
}

func TestRTBox_UnpackNegativeTest(t *testing.T) {
	g := NewGomegaWithT(t)

	testBox := NewBox(errorMapFunc, 5)
	unBoxed, err := testBox.Unpack()

	g.Expect(unBoxed).Should(BeNil())
	g.Expect(err).Should(Equal(mapErr))
}
