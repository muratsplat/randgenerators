package randgenerators

import (
	"testing"
)

func TestFirst(t *testing.T) {

	var size uint = 3

	random, err := GenereteRandString(size, NUMBER)

	if err != nil {

		t.Fatal(err)
	}

	if len(random) != 3 {

		t.Fail()

	}

}

func TestNumberUpperALPHA(t *testing.T) {

	var size uint = 18

	random, err := GenereteRandString(size, NUMBER, UPPER_CASE_ALPHA)

	if err != nil {

		t.Fatal(err)
	}

	if len(random) != 18 {

		t.Fail()

	}

	randomTwo, err := GenereteRandString(size, NUMBER, UPPER_CASE_ALPHA)

	if randomTwo == random {

		t.Fail()
	}

	numberOFupper := 0

	for i := 0; i < len(randomTwo); i++ {

		if isLower(randomTwo[i]) {

			numberOFupper++

		}

	}

	if numberOFupper == 0 {

		t.Fail()
	}

}

func isLower(char byte) bool {

	var byt byte

	for i := 0; i < len(upper_case_alphas); i++ {

		byt = upper_case_alphas[i]

		if byt == char {

			return true
		}
	}

	return false
}
