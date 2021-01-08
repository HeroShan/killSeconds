package blocker

import(
	_"os"
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

func (bs *Blockers)DataLoading(format LoadData){
	formatData := format
	// for k,v := range formatData{
	// 	log.Printf(k,v)
	// }
	log.Printf("%#v\n",formatData)
	
	
}

func chunkBucket(){

}
