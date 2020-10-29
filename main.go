package main

import(
	"SecondsKill/core"
	"fmt"
)
func main(){
	kill := core.InitList(100000)
	var i,j int
	for i<500 {
		i++
		kill.GetKillOption()
		p := kill.List.Length()
		fmt.Printf("%d",p)
	}
	go func(){
		for j<500 {
			j++
			kill.GetKillOption()
			p := kill.List.Length()
			fmt.Printf("%d",p)
		}
	}()

}