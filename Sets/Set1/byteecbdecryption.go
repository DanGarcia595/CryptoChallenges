package Set1

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

func Byteecbdecryption(plaintext []byte) string {
	bs := 16
	key := RandBytes(bs)
	As := strings.Repeat("A", bs)
	var pt string
	var OGknowntext string
	var blocklen int
	fmt.Println(len(plaintext) / bs)
	if len(plaintext)%bs == 0 {
		blocklen = len(plaintext) / bs
	} else {
		blocklen = (len(plaintext) / bs) + 1
	}
	for blocknum := 0; blocknum < blocklen; blocknum++ {
		for byteindex := 0; byteindex < bs; byteindex++ {
			inputtext := As[:bs-(byteindex+1)]
			if blocknum > 0 && len(pt) >= bs {
				OGknowntext = pt[len(pt)-(bs-1):]
			} else {
				OGknowntext = inputtext + pt
			}
			inputtext += string(plaintext)
			paddedciphertext := AESECBEncryption([]byte(inputtext), string(key))
			paddedcipherblock := paddedciphertext[blocknum*bs : (blocknum+1)*bs]
			for index := 0x00; index < 0xFF; index++ {
				knowntext := OGknowntext + string(index)
				printableknowntext := ""
				guessedblock := AESECBEncryption([]byte(knowntext), string(key))
				if compareslices(guessedblock, paddedcipherblock) == true {
					for _, element := range knowntext {
						if string(element) == "\n" {
							printableknowntext += "*"
						} else {
							printableknowntext += string(element)
						}
					}
					if string(index) == "\n" {
						fmt.Println("match! :", "*", " knowntext: ", printableknowntext)
					} else {
						fmt.Println("match! :", string(index), " knowntext: ", printableknowntext)
					}
					pt += string(index)
					break
				}
			}
		}
	}
	return pt

}

func decipherECB(plaintext []byte, Apadding int, blockoffset int, key []byte, prefix []byte) string {
	bs := 16
	var pt string
	var OGknowntext string
	var blocklen int
	fmt.Println(len(plaintext) / bs)
	if len(plaintext)%bs == 0 {
		blocklen = len(plaintext) / bs
	} else {
		blocklen = ((len(plaintext) + len(prefix)) / bs) + 1
	}
	fmt.Println("Blocklen is:", blocklen)
	for blocknum := blockoffset; blocknum < blocklen; blocknum++ {
		for byteindex := 0; byteindex < bs; byteindex++ {
			inputtext := string(prefix)
			inputtext += strings.Repeat("A", Apadding+bs-(byteindex+1))
			guessingBlockText := strings.Repeat("A", bs-(byteindex+1))
			if blocknum > 0 && len(pt) >= bs {
				OGknowntext = pt[len(pt)-(bs-1):]
			} else {
				OGknowntext = guessingBlockText + pt
			}
			inputtext += string(plaintext)
			paddedciphertext := AESECBEncryption([]byte(inputtext), string(key))
			paddedcipherblock := paddedciphertext[blocknum*bs : (blocknum+1)*bs]
			for index := 0x00; index < 0xFF; index++ {
				knowntext := OGknowntext + string(index)
				printableknowntext := ""
				guessedblock := AESECBEncryption([]byte(knowntext), string(key))
				if compareslices(guessedblock, paddedcipherblock) == true {
					for _, element := range knowntext {
						if string(element) == "\n" {
							printableknowntext += "*"
						} else {
							printableknowntext += string(element)
						}
					}
					if string(index) == "\n" {
						fmt.Println("match! :", "*", " knowntext: ", printableknowntext)
					} else {
						fmt.Println("match! :", string(index), " knowntext: ", printableknowntext)
					}
					pt += string(index)
					break
				}
			}
		}
	}
	return pt
}

func compareslices(a, b []byte) bool {
	for index, element := range b {
		if a[index] != element {
			return false
		}
	}
	return true
}

func Byteecbdecryptionharder(plaintext []byte) string {
	//Write a block size detector
	key := RandBytes(16)
	prefix := RandBytes(rand.Intn(100))
	fmt.Printf("prefix length is: %d\n", len(prefix))
	injectionText := "A"
	inputtext := string(prefix) + injectionText + string(plaintext)
	ciphertext := AESECBEncryption([]byte(inputtext), string(key))
	for !DetectECB(ciphertext) {
		injectionText += "A"
		inputtext = string(prefix) + injectionText + string(plaintext)
		//fmt.Println(injectionText)
		ciphertext = AESECBEncryption([]byte(inputtext), string(key))
	}
	bs := 16
	fmt.Println("blocksize is:", bs)
	// Find the index of the blocks you injected (They are the first repeating blocks from the cipher)
	blockoffset := Findrepeatingblocks(ciphertext, bs) - 1
	Apadding := len(injectionText) % 16
	fmt.Println("Block repeat found at: ", blockoffset, " As used: ", len(injectionText))
	fmt.Println("Prefix + Apadding % 16 = ", (len(prefix)+Apadding)%16)
	return decipherECB(plaintext, Apadding, blockoffset, key, prefix)

}

func Findrepeatingblocks(ciphertext []byte, bs int) (blockIndex int) {
	var block, blockToCompare []byte
	for i := 0; i < len(ciphertext)/bs; i++ {
		block = ciphertext[(i)*bs : (i+1)*bs]
		for j := 0; j < len(ciphertext)/bs; j++ {
			if j == i {
				continue
			}
			blockToCompare = ciphertext[j*bs : (j+1)*bs]
			if bytes.Equal(block, blockToCompare) {
				return i
			}
		}
	}
	return -1
}
