package blocker

import(
	"os"
	"log"
	_"sync"
)

type Blockers struct{
	Bmap	*map[int]blocker
	Num	int
	Status	string	
	Max	int
	FormatDataFunc	func()
}

type blocker struct{
	Index	int
	Value	interface{}
	Next 	*blocker
}

func LoadJson(path string) []byte {
	jsonData := make([]byte,1000)
	file,err := os.Open(path); if err != nil{
		log.Printf("open file is fail\n")
	}
	file.Read(jsonData)
	file.Close()
	return jsonData
}

type LoadData func(data []byte) []LoadDataStr
type LoadDataStr struct{
	Index int
	Value interface{}
}

func NewBlockers(blockerMax int) *Blockers{
	var b blocker
	b.Index = 0
	b.Value = nil
	b.Next  = &b
	bs := new(Blockers)
	bmap := make(map[int]blocker)
	bmap[0] = b
	bs.Bmap = &bmap
	bs.Num = 1
	bs.Status = "null"
	bs.Max = blockerMax
	return bs
	 
}

func (bs *Blockers)DataLoading(data []LoadDataStr){

	
	
}

func chunkBucket(){

}
