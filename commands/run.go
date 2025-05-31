package commands

import (
	"fmt"

	"github.com/MohakGupta2004/Morphsploit/module/csrf"
)

func Run(moduleName string) error {
	switch moduleName {
	case "csrf":
		csrf.Run()
	default:
		return fmt.Errorf("No run implementation found for module %s", moduleName)
	}
	return nil
}
