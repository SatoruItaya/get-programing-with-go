package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	plainText := "your message goes here"
	plainText = strings.Replace(plainText, " ", "", -1)
	plainText = strings.ToUpper(plainText)
	
	keyword := "GOLANG"
	keyIndex := 0
	message := ""
	
	for _,c := range plainText{
		c -= 'A'
		k,_ := utf8.DecodeRuneInString(string(keyword[keyIndex]))
		k -= 'A'
		
		c = (c+k+26)%26 + 'A'
		message += string(c)
		
		keyIndex++
		keyIndex %= len(keyword)
	}
	
	fmt.Println(message)
}
