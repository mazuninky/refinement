package blood_contracts_go

import (
	. "github.com/onsi/gomega"
	"regexp"
	"testing"
)

var numberRegex = regexp.MustCompile("[0-9]+")

func TestRegexMapFuncWithInvalidType(t *testing.T) {
	g := NewGomegaWithT(t)

	regexFunc := createRegexMapFunc(numberRegex)

	_, err := regexFunc(5)

	g.Expect(err).Should(HaveOccurred())
}

func TestRegexMapFunc(t *testing.T) {
	g := NewGomegaWithT(t)

	regexFunc := createRegexMapFunc(numberRegex)
	numberString := "1234567890"

	actualValue, err := regexFunc(numberString)

	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(numberString).Should(Equal(actualValue))
}