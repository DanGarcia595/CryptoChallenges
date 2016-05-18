package Set1

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func Countrepeatingbytes(ciphertext []byte) (numRepeat int) {
	bs := 16 //block size
	var sum int = 0
	var block, blockToCompare []byte
	for i := 1; i < len(ciphertext)/bs; i++ {
		if i != 0 {
			block = ciphertext[(i-1)*bs : i*bs]
		} else {
			block = ciphertext[:bs]
		}
		for j := 1; j < len(ciphertext)/bs; j++ {
			if j != 0 {
				blockToCompare = ciphertext[(j-1)*bs : j*bs]
			} else {
				blockToCompare = ciphertext[:bs]
			}
			for index, _ := range block {
				sum += int(block[index] ^ blockToCompare[index])
			}
			if sum == 0 {
				numRepeat++
			}
			sum = 0
		}
		numRepeat--
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
		tmpscore = Countrepeatingbytes(hexBytes)
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

func DetectECB(ciphertext []byte) (ret int) {
	score := Countrepeatingbytes(ciphertext)
	if score > 0 {
		fmt.Println("ECB detected. Score: ", score)
		ret = 1
	} else {
		fmt.Println("No ECB found. Most likely CBC. Score: ", score)
		ret = 0
	}
	return
}
