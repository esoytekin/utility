package utility

import "github.com/manifoldco/promptui"

func SelectPrompt(items []string, label string) int {

	prompt := promptui.Select{Label: label, Items: items}

	i, _, e := prompt.Run()

	if e != nil {
		return -1
	}

	return i
}

func ConfirmPrompt(label string) bool {

	prompt := promptui.Prompt{Label: label, IsConfirm: true}

	_, err := prompt.Run()

	if err != nil {
		return false
	}

	return true

}

func PasswordPrompt(label string, validate func(input string) error) (string, error) {

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Mask:     '*',
	}

	return prompt.Run()

}
