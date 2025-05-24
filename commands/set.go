package commands

import (
	"fmt"

	"github.com/MohakGupta2004/Morphsploit/module/csrf"
)

func Set(value string, moduleName string) {
	switch moduleName {
	case "csrf":
		csrf.Set(value)
	default:
		fmt.Println("No options by this name is found")
	}
}
