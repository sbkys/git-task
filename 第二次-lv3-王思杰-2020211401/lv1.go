package main

import "fmt"

type up struct {
	 name string
	 attention string
}
type content struct{
	 title string
	 dianzan string
	 toubi   string
	 shoucang string
	 jieshao  string
	 recommend []string
	 up
}

func main(){
     b1 := up{
     	name: "蒸汽火鸡",
     	attention: "30.6万",
     	}
     b2 :=content{
		 title:     "AMD重回巅峰",
		 dianzan:   "7.1万",
		 toubi:     "3.8万",
		 shoucang:  "1.3万",
		 jieshao:   "看完发布会直播，\n通宵做完了这期，\n值得。",
		 recommend: []string{"AMD你不讲武德", "5分钟看AMD", "AMD6000系显卡发布"},
	 }
    fmt.Println(b1)
     	fmt.Println(b2)

	 }
