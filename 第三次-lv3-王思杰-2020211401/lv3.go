package main

import "fmt"

func main() {
over := make(chan bool,10)   
for i := 0; i < 10; i++ {     //for 循环完了gorutinue还没开始
go func(i int) {
fmt.Println(i)
}(i)
if i == 9 {
over <- true
}                  //无缓冲通道在一个进程中传入一个值直接卡死不会进入下一步
}
go func(){              
	<-over
}()
fmt.Println("over!!!")
}