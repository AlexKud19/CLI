package main

import (
	"cli/app/storage"
	"fmt"
	"time"
)

func main() {
	binList := storage.NewBinList()
	bin, err := binList.NewBin("hello", true, time.Now(), "bye")
	if err != nil {
		fmt.Println(err)
		return
	}
	binList.AddBin(*bin)
	binList.ReadBinList("data.json")
}
