package server

import (
	"SecondsKill/core"
)
type coreKs interface{}
type Kschan struct{
	kspush chan<- bool
	kspop  <-chan bool
}
var over bool 
func killOptionPush(kill *core.Kill,ks Kschan){
	 over = kill.GetKillOption()
	 ks.kspush<- over
}


func killOptionPop(kill *core.Kill,ks Kschan) string{

	var writeMsg string
	over = <-ks.kspop
	if over == true{
		writeMsg = "秒杀结束"
	}else{
		writeMsg = "秒杀成功"
	}
	return writeMsg
}




