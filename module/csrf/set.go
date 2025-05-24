package csrf

import "fmt"

func Set(value string) {
	switch value {
	case "1":
		NoDefense()
	default:
		fmt.Println("Option doesn't exist")
	}
}
