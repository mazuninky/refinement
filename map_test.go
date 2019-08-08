package blood_contracts_go

import (
	. "github.com/onsi/gomega"
	"regexp"
	"testing"
)

func TestRegexMapFuncWithInvalidType(t *testing.T) {
	g := NewGomegaWithT(t)

	regex := regexp.MustCompile("[0-9]+")

	regexFunc := createRegexMapFunc(regex)

	_, err := regexFunc(5)

	g.Expect(err).Should(HaveOccurred())
}