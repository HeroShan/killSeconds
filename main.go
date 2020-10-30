package main

import(
	"SecondsKill/core"
	"fmt"
	"os"
	"time"
)
func main(){
	kill := core.InitList(10000)
	kill1 := core.InitList(10000)
	kill2 := core.InitList(20000)
	var j,p,p1,p2 int

	go func(kill *core.Kill,j int,p int){
		for j<200000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill1,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<100000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill2,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<100000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill2,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<100000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill2,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<100000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill2,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<200000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill,j,p)
	go func(kill *core.Kill,j int,p int){
		for j<200000000 {
			j++
			kill.GetKillOption()
		}
	
	}(kill,j,p)

	for{
		time.Sleep(3)
		p = kill.List.Length()
		p1 = kill1.List.Length()
		p2 = kill2.List.Length()
		if p2 == 0{
		 os.Exit(0)
		}
		fmt.Printf("i:%d ",p)
		fmt.Printf("i1:%d ",p1)
		fmt.Printf("i2:%d \n",p2)
	}
	
}
