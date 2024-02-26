package main

import (
	"fmt"
	"os"
)

func getMDPath() (string, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		return "", fmt.Errorf("No path provided")
	}
	return args[0], nil
}

func getMDContentFrom(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	str := ""
	buf := make([]byte, 100)
	for {
		n, err := file.Read(buf)
		if n > 0 {
			str += string(buf[:n])
		}
		if err != nil {
			break
		}
	}

	return str, nil
}
