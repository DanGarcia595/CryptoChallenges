package main

import (
	"fmt"
	"github.com/DanGarcia595/CryptoChallenges/Set1"
)

func main() {
	//fmt.Println(Set1.Fixedxor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))

	//fmt.Println(Set1.DecryptSinglebytexor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))

	fmt.Println(Set1.DetectSinglebytexor("/home/dan/data.txt"))
}
