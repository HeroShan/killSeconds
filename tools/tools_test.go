package tools

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"unsafe"
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

func TestGoogleAuth_VerifyCode(t *testing.T) {
	g := new(GoogleAuth)
	src := g.base32encode([]byte("c415090fd839373b673561f40c740adf5451be4a"))
	fmt.Println(src)
	b,err := g.VerifyCode(src,"054699")

	if err != nil{
		panic(err)
	}
	fmt.Printf("src:%t\n",b)
}