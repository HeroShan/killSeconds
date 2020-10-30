package main

import(
	"SecondsKill/core"
	"fmt"
	"os"
	"time"
)
func main(){
	kill := core.InitList(50000)
	var j,p int

	go func(kill *core.Kill,j int,p int){
		for j<2000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<1000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<2000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<2000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill,j,p)

	for{
		time.Sleep(3)
		p = kill.List.Length()
		if p == 0{
		 os.Exit(0)
		}
		fmt.Printf("i:%d \n",p)
	}
	
}
