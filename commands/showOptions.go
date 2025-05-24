package commands

import (
	"fmt"
	"strings"

	"github.com/MohakGupta2004/Morphsploit/module/csrf"
)

func Opt(moduleName string) {
	moduleName = strings.TrimSpace(moduleName) // ← Clean input
	fmt.Println("--------------Showing options regarding to the module--------------")
	switch moduleName {
	case "csrf":
		csrf.Options()
	default:
		fmt.Println("No options found")
	}
}
