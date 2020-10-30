package tools
//自旋锁
import(
	"sync"
	"runtime"
	"sync/atomic"
)

type spinLock uint64
func (sl *spinLock) Lock() {
    for !atomic.CompareAndSwapUint64((*uint64)(sl), 0, 1) {
        runtime.Gosched()
    }
}
func (sl *spinLock) Unlock() {
    atomic.StoreUint64((*uint64)(sl), 0)
}
func NewSpinLock() sync.Locker {
    var lock spinLock
    return &lock
}