package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type StorageDb struct {
	fileName string
}

func NewStorageDb(name string) *StorageDb {
	return &StorageDb{
		fileName: name,
	}
}

func (db *StorageDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	return data, nil
}

func (db *StorageDb) CheckJSON() bool {
	ext := filepath.Ext(db.fileName)
	return strings.EqualFold(ext, ".json")
}

func (db *StorageDb) Write(content []byte) {
	file, err := os.Create(db.fileName)
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
