package blocker

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
type Options interface{}
type LoadData func(data []byte) LoadDataStr
type LoadDataInfo struct{
	Index int
	Value interface{}
}
type LoadDataStr struct{
	Data []LoadDataInfo
	Op	 *Options
}

func NewBlockers(blockerMax int) *Blockers{
	var(
		b 	blocker
		bs 	Blockers
		bmap  []blocker
	)
	b = blocker{
		Index : 0,
		Value : nil,
		Next  : &b,
	}
	bmap = append(bmap,b)
	bs.Num = 1
	bs.Status = "null"
	bs.Max = blockerMax
	bs = Blockers{
		Bmap 	: bmap,
		Num  	: 1,
		Status	: "null",
		Max		: blockerMax,
	}
	return &bs
	 
}

func (bs *Blockers)DataLoading(format LoadData,data []byte){
	formatData := format(data)
	var(
		temp 	blocker
		b 	 	blocker
		mapLen,blockeronMapIndex  int
	)
	mapLen = bs.Num-1
	for _,v := range formatData.Data{
		if blockeronMapIndex < bs.Max{
			//头部链表追加
			temp.Index = blockeronMapIndex
			temp.Value = v
			temp.Next = &bs.Bmap[mapLen]
			bs.Bmap[mapLen] = temp
			blockeronMapIndex++
		}else{
			mapLen++
			//指针操作map 分片赋值
			//桶满分配下一个空桶
			bs.Bmap = append(bs.Bmap,b)
			blockeronMapIndex = 0
		}
	}
	formatData.Data = nil	
}

