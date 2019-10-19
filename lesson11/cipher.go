package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "your message goes here"
	plainText = strings.Replace(plainText, " ", "", -1)
	plainText = strings.ToUpper(plainText)
	
	keyword := "GOLANG"
	keyIndex := 0
	message := ""
	
	for i:=0; i<len(plainText); i++{
		c := plainText[i] - 'A'
		k := keyword[keyIndex] - 'A'
		
		c = (c+k)%26 + 'A'
		message += string(c)
		
		keyIndex++
		keyIndex %= len(keyword)
	}
	
	fmt.Println(message)
}
