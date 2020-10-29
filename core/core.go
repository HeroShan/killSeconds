package core

import(
	"SecondsKill/tools"
	"sync"
)
type Kill struct{
	SpLock sync.Locker
	List *tools.ListNode
}
func InitList(n int)(kill *Kill){
	var(
		i int
		lock sync.Locker
		list tools.ListNode
	)
	
	for i = 0; i < n; i++{
		list.AddHeadList(i)
	}
	lock = tools.NewSpinLock()
	kill.SpLock = lock
	kill.List = &list
	return kill
	
}
func (kill *Kill)GetKillOption(){
	kill.SpLock.Lock()
	kill.List.DelNode()
	kill.SpLock.Unlock()
}
