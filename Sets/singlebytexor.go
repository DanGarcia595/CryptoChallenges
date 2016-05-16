package Set1

import (
	"encoding/hex"
)

func Singlebytexor(hexstring string, xor byte) (ret []byte) {
	hexBytes, _ := hex.DecodeString(hexstring) //Decode the hex to bytes
	final := make([]byte, len(hexBytes))       //This is the xoring part
	for index, _ := range hexBytes {
		final[index] = hexBytes[index] ^ xor
	}
	ret = final
	return
}

func DecryptSinglebytexor(hexstring string) (ret string) {
	var xor byte = 0
	var bestscore int = 0
	var tmpscore int = 0
	var tmp string
	for xor = 0; xor != 255; xor++ {
		tmp = string(Singlebytexor(hexstring, xor)[:])
		tmpscore = ScoreAscii(tmp)
		if tmpscore > bestscore {
			ret = tmp
			bestscore = tmpscore
		}
	}
	return
}

func ScoreAscii(plaintext string) (ret int) {
	ret = 0
	for _, char := range plaintext {
		char := string(char)
		switch char {
		case "e", "E":
			ret += 2
		case "t", "T":
			ret += 2
		case "a", "A":
			ret += 2
		case "o", "O":
			ret += 2
		case "i", "I":
			ret += 2
		case "n", "N":
			ret += 2
		case "s", "S":
			ret += 2
		case "h", "H":
			ret += 2
		case "r", "R":
			ret += 2
		case "d", "D":
			ret += 2
		case "l", "L":
			ret += 2
		case "u", "U":
			ret += 2
		case "b", "B", "c", "C", "f", "F", "g", "G", "j", "J", "k", "K", "m", "M", "p", "P", "q", "Q", "v", "V", "w", "W", "x", "X", "y", "Y", "z", "Z", " ", "!", "@", "'":
			ret += 1
		default:
			ret += -1
		}
	}
	return
}
