package tools
import(
	"fmt"
)
//列表
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

func (this *ListNode)AddHeadList(n int) {
	var temp ListNode
	temp.Next = nil
	temp.Val = n
	for{
		if this.Next == nil {
			this.Next = &temp
			break
		}else{
			this = this.Next
		}
	}
	
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