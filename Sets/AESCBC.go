package Set1

import (
	"crypto/aes"
	"encoding/base64"
	"io/ioutil"
)

func AESCBCDecryption(ciphertext []byte, key string, IV []byte) []byte {
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
	var previousCipherBlock []byte
	for len(ciphertext) > 0 {
		cipher.Decrypt(plaintext, ciphertext)
		decryptedBlock := plaintext[:bs]
		if i == 0 {
			for index, element := range IV {
				decryptedBlock[index] ^= element
			}
		} else {
			for index, element := range previousCipherBlock {
				decryptedBlock[index] ^= element
			}
		}
		for index, element := range decryptedBlock {
			finalplaintext[(i*bs)+index] = element
		}
		i++
		plaintext = plaintext[bs:]
		previousCipherBlock = ciphertext[:bs]
		ciphertext = ciphertext[bs:]
	}
	return finalplaintext[:len(finalplaintext)-5]
}

func AESCBCEncryption(plaintext []byte, key string, IV []byte) []byte {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	bs := 16
	if len(plaintext)%bs != 0 {
		plaintext = PKCS7padding(plaintext, bs)
	}
	i := 0
	ciphertext := make([]byte, len(plaintext))
	finalciphertext := make([]byte, len(plaintext))
	var CipherBlock []byte
	for len(ciphertext) > 0 {
		if i == 0 {
			for index, element := range IV {
				plaintext[index] ^= element
			}
		} else {
			for index, element := range CipherBlock {
				plaintext[index] ^= element
			}
		}
		cipher.Encrypt(ciphertext, plaintext)
		CipherBlock = ciphertext[:bs]
		for index, element := range CipherBlock {
			finalciphertext[(i*bs)+index] = element
		}
		i++
		plaintext = plaintext[bs:]
		ciphertext = ciphertext[bs:]
	}
	return finalciphertext
}

func DecryptAESCBC(filename string, key string) (ret string) {
	content, _ := ioutil.ReadFile(filename)
	cipherBytes, _ := base64.StdEncoding.DecodeString(string(content))
	IV := make([]byte, 16)
	for i := 0; i < 16; i++ {
		IV[i] = 0x00
	}
	plainText := AESCBCDecryption(cipherBytes, key, IV)
	ret = string(plainText)
	return
}
