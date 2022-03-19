package base62

import "errors"

const ALL_CHAR = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Uint64ToBase62(value uint64) string {
	temp := make([]byte, 0, 10)
	for value > 0 {
		pos := value % 62
		temp = append(temp, ALL_CHAR[pos])
		value = value / 62
	}
	reverse(temp)

	return string(temp)
}

func Base62ToUint(str string) (uint64, error) {
	var value uint64
	var idx uint64
	for i := 0; i < len(str); i++ {
		char := str[i]
		switch {
		case char >= '0' && char <= '9':
			idx = uint64(char - '0')
		case char >= 'A' && char <= 'Z':
			idx = uint64(char - 'A' + 10)
		case char >= 'a' && char <= 'z':
			idx = uint64(char - 'a' + 36)
		default:
			return 0, errors.New("invalid byte")
		}

		value = value*62 + idx
	}

	return value, nil
}

func reverse(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}
