package Set1

import (
	"crypto/aes"
	"encoding/base64"
	//	"fmt"
	"io/ioutil"
)

func AESECB(ciphertext []byte, key string) []byte {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	bs := 16
	if len(ciphertext)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	i := 0
	plaintext := make([]byte, len(ciphertext))
	finalplaintext := make([]byte, len(ciphertext))
	for len(ciphertext) > 0 {
		cipher.Decrypt(plaintext, ciphertext)
		ciphertext = ciphertext[bs:]
		decryptedBlock := plaintext[:bs]
		for index, element := range decryptedBlock {
			finalplaintext[(i*bs)+index] = element
		}
		i++
		plaintext = plaintext[bs:]
	}
	return finalplaintext[:len(finalplaintext)-5]
}

func DecryptAESECB(filename string, key string) (ret string) {
	content, _ := ioutil.ReadFile(filename)
	cipherBytes, _ := base64.StdEncoding.DecodeString(string(content))
	plainText := AESECB(cipherBytes, key)
	ret = string(plainText)
	return
}
