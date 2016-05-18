package main

import (
	//	"encoding/hex"
	//	"fmt"
	"github.com/DanGarcia595/CryptoChallenges/Sets"
)

func main() {
	//fmt.Println(Set1.Fixedxor("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))

	//fmt.Println(Set1.DecryptSinglebytexor("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))

	//fmt.Println(Set1.DetectSinglebytexor("/home/dan/data.txt"))

	//	cipherText, _ := hex.DecodeString(Set1.Repeatingkeyxor(`Burning 'em, if you ain't quick and nimble
	//I go crazy when I hear a cymbal`, "I Love Coke"))

	//	fmt.Println(string(cipherText))
	//	plainText, _ := hex.DecodeString(Set1.Repeatingkeyxor(string(cipherText), "I Love Coke"))
	//	fmt.Println(string(plainText))

	//fmt.Println(Set1.Hammingdistance([]byte("this is a test"), []byte("wokka wokka!!!")))

	//fmt.Println(Set1.Breakrepeatingkeyxor("/home/dan/6.txt"))

	//fmt.Println(Set1.DecryptAESECB("/home/dan/7.txt", "YELLOW SUBMARINE"))

	//	lineNumber, lineValue := Set1.DetectECBFromFile("/home/dan/8.txt")
	//	fmt.Println("Best Guess: ", lineNumber, " with value", lineValue)

	//fmt.Println(string(Set1.PKCS7padding([]byte("DAN"), 5)))

	//IV := make([]byte, 16)

	//fmt.Println(Set1.DecryptAESCBC("/home/dan/10.txt", "YELLOW SUBMARINE"))
	plaintext := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	ciphertext := Set1.Encryption_oracle([]byte(plaintext))
	Set1.DetectECB(ciphertext)
	//fmt.Println(hex.EncodeToString(ciphertext))
	//ciphertext = Set1.AESECBEncryption([]byte(plaintext), "Dan is the best.")
	//foo = Set1.DetectECB(ciphertext)
	//fmt.Println(hex.EncodeToString(ciphertext))
	//thenewplaintext := Set1.AESECBDecryption(ciphertext, "Dan is the best.")

}
