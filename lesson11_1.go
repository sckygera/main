package main

import (
	"fmt"
	"strings"
)

func main() {

	stri1 := strings.SplitAfterN("Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software.", " ", 18)
	capitalLetter := 0
	for _, item := range stri1 {

		v := strings.ToLower(item)
		s := strings.Compare(v, item)
		if s != 0 {
			capitalLetter++
		}

	}
	fmt.Println("Строка содержит", capitalLetter, "слов с большой буквы")
	fmt.Sprintf("%s:%d", filename, f.Line)

}
