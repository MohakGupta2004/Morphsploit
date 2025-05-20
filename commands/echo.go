package commands

import (
	"fmt"
	"strings"
)

func Echo(txt []string) {
	value := strings.Join(txt, " ")
	fmt.Printf("echo: %s\n", value)
}
