package main

import (
	"fmt"

	"github.com/esoytekin/utility"
)

func testPrompt() {
	label := "select a fruit"
	items := []string{"apple", "orange"}

	index := utility.SelectPrompt(items, label)

	if index != -1 {
		fmt.Printf("you selected %s\n", items[index])
	} else {
		fmt.Println("canceled")
	}

}

func main() {
	testPrompt()
}
