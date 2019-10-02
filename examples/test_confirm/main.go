package main

import (
	"fmt"

	"github.com/esoytekin/utility"
)

func main() {

	sure := utility.ConfirmPrompt("Are you sure?")

	if sure {
		fmt.Println("yes, you are!")
	} else {
		fmt.Println("well, you have doubts...")
	}
}
