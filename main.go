package main

import(
	"SecondsKill/core"
	"fmt"
	"os"
)
func main(){
	kill := core.InitList(100)
	var j,l int
	var over bool
	go kill.Monitor()
	//俩个线程模拟秒杀
	go func(){
		for j<20000 {
			j++
			over = kill.GetKillOption()
			if over == true{
				fmt.Println("秒杀结束！")
				os.Exit(0)
			}
		}

	}()
	for j<200000 {
			j++
			over = kill.GetKillOption()
			if over == true{
				fmt.Println("秒杀结束！")
				os.Exit(0)
			}
			l = kill.List.Length()
			fmt.Printf("还剩%d个\n",l)
		}
	
	
	


	
}
