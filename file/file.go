package file

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}

func CheckJSON(name string) (bool, error) {
	data, err := ReadFile(name)
	if err != nil {
		return false, err
	}
	return json.Valid(data), nil
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
