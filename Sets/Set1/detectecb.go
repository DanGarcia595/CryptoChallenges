package Set1

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func Countrepeatingblocks(ciphertext []byte) (numRepeat int) {
	bs := 16 //block size
	var block, blockToCompare []byte
	for i := 0; i < len(ciphertext)/bs; i++ {
		block = ciphertext[(i)*bs : (i+1)*bs]
		for j := 0; j < len(ciphertext)/bs; j++ {
			if j == i {
				continue
			}
			blockToCompare = ciphertext[j*bs : (j+1)*bs]
			if bytes.Equal(block, blockToCompare) {
				numRepeat++
			}
		}
	}
	return
}

func DetectECBFromFile(filename string) (lineNumber int, ret string) {
	var bestscore int = 0
	var tmpscore int = 0
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var i int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexBytes, _ := hex.DecodeString(scanner.Text())
		tmpscore = Countrepeatingblocks(hexBytes)
		if tmpscore > 2 {
			fmt.Println("ECB possible on line ", i, " with ", tmpscore-1, " matching blocks")
		}
		if tmpscore > bestscore {
			ret = scanner.Text()
			lineNumber = i
			bestscore = tmpscore
		}
		i++
	}
	return
}

func DetectECB(ciphertext []byte) (ret bool) {
	score := Countrepeatingblocks(ciphertext)
	if score > 0 {
		//fmt.Println("ECB detected. Score: ", score)
		ret = true
	} else {
		//fmt.Println("No ECB found. Most likely CBC. Score: ", score)
		ret = false
	}
	return
}
