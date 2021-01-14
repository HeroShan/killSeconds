package blocker

import(
	_"os"
	"log"
	_"sync"
)

type Blockers struct{
	Bmap	[]blocker
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
	b.Next  = nil
	bs := new(Blockers)
	var bmap  []blocker
	bmap = append(bmap,b)
	bs.Bmap = bmap
	bs.Num = 1
	bs.Status = "null"
	bs.Max = blockerMax
	return bs
	 
}

func (bs *Blockers)DataLoading(format LoadData,data []byte){
	formatData := format(data)
	var(
		temp 	blocker
		b 	 	blocker
		mapLen,blockeronMapIndex  int
		blockeronMap	blocker
	)
	
	//log.Printf("%#v\n",bs.Bmap[0])
	mapB := bs.Bmap
	mapLen = bs.Num-1
	
	for k,v := range formatData.data{
		blockeronMap = bs.Bmap[mapLen]
		blockeronMapIndex = blockeronMap.Index
		if blockeronMapIndex < bs.Max{
			for blockeronMap.Next != nil{
				blockeronMap = *blockeronMap.Next
			}
			//log.Println(k)
			temp.Index = k
			temp.Value = v
			blockeronMap.Next = &temp
		}else{
			mapLen++
			//指针操作map 分片赋值
			mapB = append(mapB,b)
		}
	}

	log.Printf("%#v\n",len(bs.Bmap))
	
	
}

func chunkBucket(){

}
