package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNSROCNKFD"
	keyword := "GOLANG"
	keyIndex := 0
	message := ""
	
	for _,c := range cipherText{
		c -= 'A'
		k,_ := utf8.DecodeRuneInString(string(keyword[keyIndex]))
		k -= 'A'
		
		c = (c-k+26)%26 + 'A'
		message += string(c)
		
		keyIndex++
		keyIndex %= len(keyword)
	}
	
	fmt.Println(message)
}
