package Set1

import (
	"bufio"
	"log"
	"os"
)

func DecryptSinglebytexor2(hexstring string) (ret string, bestscore int) {
	var xor byte = 0
	var tmpscore int = 0
	var tmp string
	for xor = 0; xor != 255; xor++ {
		tmp = string(Singlebytexor(hexstring, xor)[:])
		tmpscore = ScoreAscii(tmp)
		if tmpscore > bestscore {
			ret = tmp
			bestscore = tmpscore
		}
	}
	return
}

func DetectSinglebytexor(filename string) (ret string) {
	var bestscore int = 0
	var tmpscore int = 0
	var tmp string
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tmp, tmpscore = DecryptSinglebytexor2(scanner.Text())
		if tmpscore > bestscore {
			ret = tmp
			bestscore = tmpscore
		}
	}
	return
}
