package bins

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Bin struct {
	Id        string    `json:"id,omitempty"`
	Private   bool      `json:"private,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name,omitempty"`
}

func (*BinList) NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}, nil
}

type BinList []Bin

func NewBinList() BinList {
	binList := make(BinList, 0, 10)
	return binList
}

func (list *BinList) AddBin(bin Bin) {
	*list = append(*list, bin)
}

func (list *BinList) ToBytes() ([]byte, error) {
	return json.Marshal(*list)
}

// func (list *BinList) save() {
// 	data, err := list.ToBytes()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	file.Write(data)
// }

func (list *BinList) ReadBinList(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
