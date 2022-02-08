package Set1

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandBytes(n int) []byte {
	b := make([]byte, n)
	for i, _ := range b {
		b[i] = byte(rand.Intn(256))
	}
	return b
}

func Encryption_oracle(plaintext []byte) []byte {
	paddings := (5 + rand.Intn(5))
	realplaintext := make([]byte, len(plaintext)+2*paddings)
	for i := 0; i < paddings; i++ {
		realplaintext[i] = 0x04
		realplaintext[i+len(plaintext)+paddings] = 0x04
	}
	for index, element := range plaintext {
		realplaintext[index+paddings] = element
	}
	choice := rand.Intn(2)
	key := RandBytes(16)
	var ciphertext []byte
	if choice == 1 { //cbc aes-128
		IV := RandBytes(16)
		fmt.Println("CBC chosen")
		ciphertext = AESCBCEncryption(realplaintext, string(key), IV)
	} else { //ecb aes-128
		fmt.Println("ECB chosen")
		ciphertext = AESECBEncryption(realplaintext, string(key))
	}
	return ciphertext
}
