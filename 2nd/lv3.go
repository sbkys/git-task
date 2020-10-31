package main

import "fmt"


type content struct{
	title string
	dz int64
	tb int64
	sc int64
	jieshao  string
	recommend []string
}
func (v *content)dianzan(){
     v.dz++
}
func (v *content)toubi(){
	v.tb++
}
func (v *content)shoucang(){
	v.sc++
}
func (v *content)yijiansanlian(){
	v.dz++
	v.tb++
	v.sc++
}
func s ()
func main()  {
	var a content
	a.dz = 71000
	a.tb = 38000
	a.sc = 13000
	fmt.Println(a.dz,a.tb,a.sc)
	a.dianzan()
	a.toubi()
	a.shoucang()
	fmt.Println(a.dz,a.tb,a.sc)
	a.yijiansanlian()
	fmt.Println(a.dz,a.tb,a.sc)
}
