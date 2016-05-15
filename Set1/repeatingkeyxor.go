package Set1

import "encoding/hex"

func Repeatingkeyxor(plaintext string, key string) (ret string) {
	hexBytes := []byte(plaintext) //Decode the hex to bytes
	final := make([]byte, len(hexBytes))
	keyBytes := []byte(key)
	for index, _ := range hexBytes {
		final[index] = hexBytes[index] ^ keyBytes[index%3]
	}
	ret = hex.EncodeToString(final)
	return
}
