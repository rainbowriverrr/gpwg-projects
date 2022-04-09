package main

//This is for Capstone 2 (Vigenere Cipher)

import (
	"fmt"
	"strings"
)

func decipher() {
	var cipherText = "CSOITEUIWUIZNSROCNKFD"
	var keyword = []rune("GOLANG")
	var decoded string = ""
	var temp rune

	for i, c := range cipherText {

		currKey := keyword[i%len(keyword)] - 'A'
		currChar := c - 'A'

		temp = (currChar-currKey+26)%26 + 'A'

		decoded += string(temp)
	}

	fmt.Printf("Decoded: %v \n", decoded)
}

func encode() {
	plainText := "WEDIGYOULUVTHEGOPHERS "
	keyword := []rune("GOLANG")

	plainText = strings.ToUpper(plainText)
	plainText = strings.Replace(plainText, " ", "", -1)

	var temp rune
	var encoded string = ""

	for i, c := range plainText {
		currKey := keyword[i%len(keyword)] - 'A'
		currChar := c - 'A'

		temp = (currChar+currKey)%26 + 'A'

		encoded += string(temp)
	}

	fmt.Printf("Encoded: %v \n", encoded)
}

func runEncoding() {

	decipher()
	encode()

}
