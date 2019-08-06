package refinement

import (
	"errors"
	"testing"
)
import . "github.com/onsi/gomega"

func TestUnpackWithPositiveTestMapFunc(t *testing.T) {
	g := NewGomegaWithT(t)

	testValue := 5
	testMapFunc := func(value interface{}) (interface{}, error){
		return value, nil
	}

	testBox := NewBox(testMapFunc, testValue)
	unBoxed, err := testBox.Unpack()

	g.Expect(unBoxed).To(Equal(testValue))
	g.Expect(err).ShouldNot(HaveOccurred())
}


func TestUnpackWithNegativeTestMapFunc(t *testing.T) {
	g := NewGomegaWithT(t)

	negativeError := errors.New("negative")
	testMapFunc := func(value interface{}) (interface{}, error){
		return nil, negativeError
	}

	testBox := NewBox(testMapFunc, 5)
	unBoxed, err := testBox.Unpack()

	g.Expect(unBoxed).Should(BeNil())
	g.Expect(err).Should(Equal(negativeError))
}
