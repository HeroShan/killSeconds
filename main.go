package main

import(
	"SecondsKill/core"
	"fmt"
)
func main(){
	kill := core.InitList(20)
	var i,j,p int
	defer func(){
		p = kill.List.Length()
		fmt.Printf("i:%d \n",p)
	}()

	go func(kill *core.Kill,j int,p int){
		for j<5 {
			j++
			kill.GetKillOption()
			p = kill.List.Length()
			fmt.Printf("j:%d \n",p)
		}
	
	}(kill,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<5 {
			j++
			kill.GetKillOption()
			p = kill.List.Length()
			fmt.Printf("j1:%d \n",p)
		}
	
	}(kill,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<5 {
			j++
			kill.GetKillOption()
			p = kill.List.Length()
			fmt.Printf("j2:%d \n",p)
		}
	
	}(kill,j,p)
	for i<20 {
		i++
		kill.GetKillOption()
		p = kill.List.Length()
		fmt.Printf("i:%d \n",p)
	}
	
}