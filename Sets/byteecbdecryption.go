package Set1

import "fmt"

func Byteecbdecryption(plaintext []byte) string {
	bs := 16
	key := "0123456789ABCDEF"
	As := "AAAAAAAAAAAAAAAA"
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
			paddedciphertext := AESECBEncryption([]byte(inputtext), key)
			paddedcipherblock := paddedciphertext[blocknum*bs : (blocknum+1)*bs]
			for index := 0x00; index < 0xFF; index++ {
				knowntext := OGknowntext + string(index)
				printableknowntext := ""
				guessedblock := AESECBEncryption([]byte(knowntext), key)
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
