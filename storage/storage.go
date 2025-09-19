package storage

import (
	"cli/app/bins"
	"cli/app/file"
	"encoding/json"
	"fmt"
	"os"
)

func SaveBins(bins bins.BinList) {
	data, err := json.Marshal(bins)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteFile(data, "data.json")
}

func ReadBinList(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(data))
	return data, nil
}
