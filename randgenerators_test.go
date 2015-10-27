package randgenerators

import (
	"testing"
)

func TestFirst(t *testing.T) {

	var size uint = 3

	random, err := genereteRandString(size, NUMBER)

	if err != nil {

		t.Fatal(err)
	}

	if len(random) != 3 {

		t.Fail()

	}

}

func TestNumberUpperALPHA(t *testing.T) {

	var size uint = 18

	random, err := genereteRandString(size, NUMBER, UPPER_CASE_ALPHA)

	if err != nil {

		t.Fatal(err)
	}

	if len(random) != 18 {

		t.Fail()

	}

	randomTwo, err := genereteRandString(size, NUMBER, UPPER_CASE_ALPHA)

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

	for i := 0; i < len(UPPER_CASE_ALPHAS); i++ {

		byt = UPPER_CASE_ALPHAS[i]

		if byt == char {

			return true
		}
	}

	return false
}
