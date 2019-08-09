package blood_contracts_go

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestRType_IsValidWithNegativeMapFunc(t *testing.T) {
	g := NewGomegaWithT(t)

	rType := NewType(negativeMapFunc)

	g.Expect(rType.IsValid(5)).Should(BeFalse())
}

func TestRType_IsValidWithPositiveMapFunc(t *testing.T) {
	g := NewGomegaWithT(t)

	rType := NewType(positiveMapFunc)

	g.Expect(rType.IsValid(5)).Should(BeTrue())
}

//func TestRType_Or(t *testing.T) {
//	g := NewGomegaWithT(t)
//
//	positiveType := NewType(positiveMapFunc)
//
//	//g.Expect(rType.IsValid(5)).Should(BeTrue())
//}