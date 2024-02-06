package main

import "fmt"
import "unicode/utf8"

func main() {

	// "" by defauld
	var str string

	var hello string = "Hello \n\t"

	// raw string
	var world string = `World \n\t `

	// support utf8 by default

	// singe quotes used for bytes (uint8)
	var rawBinary byte = '\x27'

	// rune (uint32) for full-utf8 symbols
	var someChinese rune = '卦'

	// string are immutable

	test := "helloworld"

	byteLen := len(test)
	symbolsLen := utf8.RuneCountInString(test)

	part1 := test[:6] // 0-5 bytes, not symbols
	part2 := test[0]  // byte 104, not "h"

	// slice byte to string and vice versa
	byteString := []byte(test)
	stringString := string(byteString)

	fmt.Println(str, hello, world, rawBinary, someChinese, test, byteLen, symbolsLen, part1, part2, byteString, stringString)
}
