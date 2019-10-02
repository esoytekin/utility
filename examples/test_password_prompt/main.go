package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/esoytekin/utility"
)

func main() {
	r, err := utility.PasswordPrompt("enter password", func(input string) error {

		if len(input) < 6 {
			return errors.New("input can't be shorter than 6 characters!")
		}

		return nil

	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("you entered: %s\n", r)

}
