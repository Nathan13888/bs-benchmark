package main

import (
	"os"
	"time"
)

func GetShortTime() string {
	return time.Now().Format("20060102150405")
}

func Mkdir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// create directory
		err := os.Mkdir(path, 0755)

		// TODO: check if error is "file exists"
		return err
	}

	// path already exists
	return nil
}
