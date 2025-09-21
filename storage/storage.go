package storage

import (
	"cli/app/bins"
	"encoding/json"
	"fmt"
	"os"
)

type StorageDb interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Storage struct {
	BinList bins.BinList
	db      StorageDb
}

func NewStorage(db StorageDb) *Storage {
	file, err := db.Read()
	if err != nil {
		return &Storage{
			BinList: bins.BinList{},
			db:      db,
		}
	}
	var binlist bins.BinList
	err = json.Unmarshal(file, &binlist)
	if err != nil {
		fmt.Println(err)
		return &Storage{
			BinList: bins.BinList{},
			db:      db,
		}
	}
	fmt.Println(23, binlist)
	return &Storage{
		BinList: binlist,
		db:      db,
	}
}

func (storage *Storage) ToBytes() ([]byte, error) {
	return json.Marshal(storage.BinList)
}

func (storage *Storage) SaveBins() {
	data, err := storage.ToBytes()
	if err != nil {
		fmt.Println(err)
	}
	storage.db.Write(data)
}

func (storage *Storage) ReadBinList(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	return data, nil
}

func (storage *Storage) AddBin(bin *bins.Bin) {
	fmt.Println(1, storage.BinList)
	storage.BinList = append(storage.BinList, *bin)
	fmt.Println(storage.BinList)
	storage.SaveBins()
}
