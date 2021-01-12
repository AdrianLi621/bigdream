package main

import "bigdream/huigou/pkg"

func main()  {
	ch:=make(chan int)
	for i:=0;i<100;i++{
		go func() {
			pkg.CusToWorks("rereir")
		}()
	}
	<-ch
}






