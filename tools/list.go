package tools
import(
	"fmt"
)
//链表
type ListNode struct {
    Val 	int
  	Next 	*ListNode
}

func (this *ListNode)AddTailList(n int) {
	var temp ListNode
	
		for this.Next != nil{
			this = this.Next
		}
		temp.Val = n
		temp.Next = nil
		this.Next = &temp	
			
}
//头部空添加
func (this *ListNode)AddHeadList(n int)(*ListNode){
	var temp ListNode
	temp.Next = this
	temp.Val = n
	return &temp	
}


func (node *ListNode)DelNode() {
	if(node.Next==nil){
		node = nil
		return
	}
    next := node.Next
    for next.Next != nil {
        node.Val = next.Val
        node = next
        next = node.Next
    }
    node.Val = next.Val
    node.Next = nil
}

//计算链表长度
func (this *ListNode)Length() (n int) {
	for this.Next != nil{
		this = this.Next
		n++
	}
	return n
}


func PrintList(this *ListNode) {
	
	for{
		if this.Next != nil {
			fmt.Println(this.Val)
			this = this.Next
		}else{
			fmt.Println(this.Val)
			break
		}
	}
}