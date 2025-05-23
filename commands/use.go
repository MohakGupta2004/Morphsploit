package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Use(args []string) error {
	pathName := strings.Split(args[0], "/")
	err, moduleName := extractModule(pathName)
	if err != nil {
		return err
	}
	println(moduleName)
	return nil
}

func extractModule(path []string) (error, string) {
	moduleName := path[1]
	if moduleName != "module" {
		return errors.New("Wrong keyword specified, Please specify module name"), "Module not Found"
	}
	info, err := os.Stat(fmt.Sprintf("./module/%s", path[2]))
	if os.IsNotExist(err) {
		return errors.New("Module not found"), "Module not found"
	}

	if !info.IsDir() {
		return errors.New("No directory found"), "No directory found"
	}

	return nil, path[2]
}
