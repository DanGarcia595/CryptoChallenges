package Set1

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func Hammingdistance(firstBlock, secondBlock []byte) (numBits int) {
	var tmpByte byte
	for index, _ := range firstBlock {
		for tmpByte = firstBlock[index] ^ secondBlock[index]; tmpByte > 0; tmpByte >>= 1 {
			numBits += int(tmpByte & byte(1))
		}
	}

	return
}

func Breakrepeatingkeyxor(filename string) (ret string) {
	content, _ := ioutil.ReadFile(filename)
	cipherBytes, _ := base64.StdEncoding.DecodeString(string(content))
	ret = Decryptrepeatingkeyxor(cipherBytes)
	hexBytes, _ := hex.DecodeString(ret)
	ret = string(hexBytes)
	return
}

func Decryptrepeatingkeyxor(cipherBytes []byte) (ret string) {
	var bestdistance int = 10000
	var finalkeysize int
	for keysize := 2; keysize <= 40; keysize++ {
		distance1 := Hammingdistance(cipherBytes[0:keysize], cipherBytes[keysize:(2*keysize)])
		distance2 := Hammingdistance(cipherBytes[(2*keysize):(3*keysize)], cipherBytes[3*keysize:(4*keysize)])
		distance4 := Hammingdistance(cipherBytes[keysize:(2*keysize)], cipherBytes[3*keysize:(4*keysize)])
		distance := (distance1 + distance2 + distance4) / 3
		distance /= keysize
		if distance <= bestdistance {
			bestdistance = distance
			finalkeysize = keysize
		}
		//	fmt.Println("Distance is", distance)
	} //By here we should know how long the key is

	fmt.Println("Keysize is", finalkeysize)
	numBlocks := finalkeysize
	key := make([]byte, finalkeysize)
	var cipherBlockLength int
	if len(cipherBytes)%finalkeysize == 0 {
		cipherBlockLength = int(len(cipherBytes) / finalkeysize)
	} else {
		cipherBlockLength = int(len(cipherBytes)/finalkeysize) + 1
	}

	for j := 0; j < numBlocks; j++ {
		block := make([]byte, cipherBlockLength)
		for i := 0; i < cipherBlockLength; i++ {
			if (finalkeysize*i)+j < len(cipherBytes) { //error handling for out of range
				block[i] = cipherBytes[(finalkeysize*i)+j]
			} else {
				block[i] = 0
			}
		}
		//We have our cipher block to decrypt
		hexBlockString := hex.EncodeToString(block)
		_, blockKey := DecryptSinglebytexor3(hexBlockString)
		key[j] = blockKey
	}
	fmt.Println("Key is", string(key))
	//	fmt.Println("cipherText is", string(cipherBytes))
	ret = Repeatingkeyxor(string(cipherBytes), string(key))
	return

}

func DecryptSinglebytexor3(hexstring string) (ret string, key byte) {
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
			key = xor
		}
	}
	return
}
