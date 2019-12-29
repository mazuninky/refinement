package _old

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRTBox_UnpackPositiveTestMapFunc(t *testing.T) {
	g := NewGomegaWithT(t)

	testValue := 5

	testBox := NewBox(positiveMapFunc, testValue)
	unBoxed, err := testBox.Unpack()

	g.Expect(unBoxed).To(Equal(testValue))
	g.Expect(err).ShouldNot(HaveOccurred())
}

func TestRTBox_UnpackWithNegativeTestMapFunc(t *testing.T) {
	g := NewGomegaWithT(t)

	testBox := NewBox(negativeMapFunc, 5)
	unBoxed, err := testBox.Unpack()

	g.Expect(unBoxed).Should(BeNil())
	g.Expect(err).Should(Equal(negativeMapErr))
}
