package Set1

import (
	"encoding/hex"
)

func Fixedxor(first, second string) (ret string) {
	firstBytes, _ := hex.DecodeString(first)   //Decode the hex to bytes
	secondBytes, _ := hex.DecodeString(second) //Decode the hex to bytes
	final := make([]byte, len(firstBytes))     //This is the xoring part
	for index, _ := range firstBytes {
		final[index] = firstBytes[index] ^ secondBytes[index]
	}
	ret = hex.EncodeToString(final)
	return
}
