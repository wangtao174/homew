package main

import (
	"fmt"
	"sync"
)

var wg   sync.WaitGroup
var count int

func judge(n int ){
	flag :=true
	for i:=2;i<n/2;i++{
		if n%i==0{
			flag=false
			break
		}
	}
	if flag{
		count++
		fmt.Println(n)
	}

	wg.Done()
}

func main() {

	for i:=2;i<50000;i++{
		wg.Add(1)
		 go judge(i)

	}
	fmt.Println(count)
	wg.Wait()
	fmt.Println(count)
	//for range ans{
	//	fmt.Println(ans)
	//}
}
