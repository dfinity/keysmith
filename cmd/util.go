package cmd

import (
	"fmt"
	"io/fs"
	"os"
)

func writeFileOrStdout(file string, output []byte, perm fs.FileMode) error {
	if file == "-" {
		fmt.Print(string(output))
		return nil
	} else {
		handle, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_EXCL, perm)
		if err != nil {
			return fmt.Errorf("Output file already exists: %s", file)
		}
		defer handle.Close()
		n, err := handle.Write(output)
		if err != nil || n != len(output) {
			return fmt.Errorf("Cannot write output to file: %s", file)
		}
		return nil
	}
}
