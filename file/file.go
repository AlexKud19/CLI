package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}

func CheckJSON(name string) bool {
	ext := filepath.Ext(name)
	return strings.EqualFold(ext, ".json")
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
