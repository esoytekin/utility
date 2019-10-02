package utility

import "github.com/manifoldco/promptui"

func AskPrompt(items []string, label string) int {

	prompt := promptui.Select{Label: label, Items: items}

	i, _, e := prompt.Run()

	if e != nil {
		return -1
	}

	return i
}
