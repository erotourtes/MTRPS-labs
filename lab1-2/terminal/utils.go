package terminal

import (
	"fmt"
	"os"
	"strings"
)

func getContentFromInput(path string) (string, error) {
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

func output(path, content string) error {
	if path == "" {
		fmt.Printf(content)
		return nil
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func mapToStr[T any](m map[string]T) string {
	s := new(strings.Builder)
	for k := range m {
		s.WriteString(fmt.Sprintf("'%s'", k))
		s.WriteString(", ")
	}
	return s.String()
}
