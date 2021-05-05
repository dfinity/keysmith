package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

func writeFileOrStdout(file string, output []byte, perm fs.FileMode) error {
	if file == "-" {
		fmt.Print(string(output))
		return nil
	} else {
		_, err := os.Stat(file)
		if !os.IsNotExist(err) {
			return fmt.Errorf(
				"Output file already exists: %s",
				file,
			)
		}
		return ioutil.WriteFile(file, output, perm)
	}
}
