package main

import (
	"cli/app/bins"
	"cli/app/storage"
	"fmt"
	"time"
)

func main() {
	binList := bins.NewBinList()
	bin, err := binList.NewBin("hellorrr", true, time.Now(), "bye")
	if err != nil {
		fmt.Println(err)
		return
	}
	binList.AddBin(*bin)
	storage.SaveBins(binList)
	storage.ReadBinList("data.json")
}
