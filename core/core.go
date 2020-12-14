package core

import(
	"SecondsKill/tools"
	"sync"
	"time"
)
type Kill struct{
	SpLock sync.Locker
	List *tools.ListNode
	stop  bool //停止秒杀

}
func InitList(n int)(*Kill){
	var i int
	list := new(tools.ListNode)
	for i = 0; i < n; i++{
		list = list.AddHeadList(i)
	}
	kill := new(Kill)
	kill.SpLock = tools.NewSpinLock()
	kill.List = list
	return kill
	
}
//监听器  监听链表为空停止秒杀
func (kill *Kill)Monitor(){
	var listLength int
	for{
		listLength = kill.List.Length()
		if listLength == 0{
			kill.stop = true
			break
		}
		time.Sleep(1/100)
		kill.stop = false
	}

}
//获取一条秒杀记录
func (kill *Kill)GetKillOption()(over bool){
	if kill.stop == true{
		return kill.stop
	}
	kill.SpLock.Lock()
	kill.List.DelNode()
	kill.SpLock.Unlock()
	return kill.stop
}
