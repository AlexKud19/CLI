package main

import (
	"cli/app/bins"
	"cli/app/file"
	"cli/app/storage"
	"fmt"
	"time"
)

func main() {
	storage := storage.NewStorage(file.NewStorageDb("data.json"))
	storage.BinList = bins.NewBinList()
	bin, err := storage.BinList.NewBin("hellorrr", true, time.Now(), "wwww")
	if err != nil {
		fmt.Println(err)
		return
	}
	storage.AddBin(bin)
	storage.BinList.ReadBinList("data.json")
}
