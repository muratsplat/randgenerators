// TODO: document this package
package randgenerators

import (
	"crypto/rand"
	"errors"
)

const (
	LOWER_CASE_ALPHAS = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	UPPER_CASE_ALPHAS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS           = "0123456789"
)

type RandType byte

const (
	NUMBER RandType = iota
	UPPER_CASE_ALPHA
	LOWER_CASE_ALPHA
)

// to generetes random characters
func GenereteRandString(strSize uint, types ...RandType) (dicts string, err error) {

	if len(types) == 0 {

		err = errors.New("At the very least a randType should be selected !")

		return
	}

	for _, v := range types {

		switch v {

		case NUMBER:
			dicts += NUMBERS
			break
		case UPPER_CASE_ALPHA:
			dicts += UPPER_CASE_ALPHAS
			break

		case LOWER_CASE_ALPHA:
			dicts += LOWER_CASE_ALPHAS
			break
		}

	}

	return makeRandom(strSize, dicts), err
}

// to mix up characters..
func makeRandom(size uint, dicts string) string {

	var bytes = make([]byte, size)

	rand.Read(bytes)

	for k, v := range bytes {
		bytes[k] = dicts[v%byte(len(dicts))]
	}

	return string(bytes)
}
