package csrf

import "fmt"

var selectedOption string

func Set(value string) {
	switch value {
	case "1", "2", "3":
		selectedOption = value
		fmt.Println("CSRF mode set to option", value)
	default:
		fmt.Println("Option doesn't exist")
	}
}

func Run() {
	switch selectedOption {
	case "1":
		NoDefense()
	default:
		fmt.Println("No valid option set")
	}
}
