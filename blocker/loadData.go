package blocker

import(
	_"os"
	_"log"
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

type LoadData func(data []byte) LoadDataStr
type LoadDataInfo struct{
	Index int
	Value interface{}
}
type LoadDataStr struct{
	data []LoadDataInfo
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

func (bs *Blockers)DataLoading(format LoadData,data []byte){
	formatData := format(data)
	var(
		
	)
	for k,_ := range formatData.data{
		k++
		//log.Printf("%v,%v\n",k,v)
	}
	//log.Printf("%#v\n",formatData)
	
	
}

func chunkBucket(){

}
