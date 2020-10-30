package core

import(
	"SecondsKill/tools"
	"sync"
)
type Kill struct{
	SpLock sync.Locker
	List *tools.ListNode
}
func InitList(n int)(*Kill){
	var(
		i int
		list tools.ListNode
	)
	
	for i = 0; i < n; i++{
		list.AddHeadList(i)
	}
	kill := new(Kill)
	kill.SpLock = tools.NewSpinLock()
	kill.List = &list
	return kill
	
}
func (kill *Kill)GetKillOption(){
	kill.SpLock.Lock()
	kill.List.DelNode()
	kill.SpLock.Unlock()
}
