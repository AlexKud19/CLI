package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}, nil
}

type BinList []Bin

func NewBinList() BinList {
	binList := make(BinList, 0, 10)
	return binList
}
