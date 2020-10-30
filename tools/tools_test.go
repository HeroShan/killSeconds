package tools

import(
	"testing"
	"fmt"
	"reflect"
	"unsafe"
	"sync"
)

func TestSpinLocks(t *testing.T){
	var p sync.Locker
	p = NewSpinLock()
	d := p
	d.Lock()
	fmt.Printf("%v\n",reflect.TypeOf(p))
	p.Unlock()
	fmt.Printf("%p\n",p)
	fmt.Printf("%v\n",unsafe.Sizeof(p))
	

}