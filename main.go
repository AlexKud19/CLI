package main

import (
	"cli/app/bins"
	"cli/app/file"
	"cli/app/storage"
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	storage := storage.NewStorage(file.NewStorageDb("data.json"))
	bin, err := bins.NewBin("hellorrr", true, time.Now(), "wwww")
	if err != nil {
		fmt.Println(err)
		return
	}
	storage.AddBin(bin)
	storage.BinList.ReadBinList("data.json")
}
