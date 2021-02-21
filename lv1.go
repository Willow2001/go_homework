package main

import "fmt"

//人类
type person struct {
	name string // 姓名
	age int // 年龄
	gender string // 性别
}
//鸽子
type dove interface {
	gu() // 鸽
}
func (p *person) gu() {
	fmt.Println(p.name, "又鸽了")
}
//复读机
type repeater interface {
	repeat(string) // 复读
}
func (p *person) repeat(word string) {
	fmt.Println(word)
}
//柠檬精
type lemoner interface {
	sour() // 酸
}
func (p *person) sour() {
	fmt.Println("我好酸啊呜呜呜")
}
//真香怪
type regreter interface {
	regret(int) // 真香
}
func (p *person) regret(i int) {
	if i>=1{
		fmt.Println("真香！")
	}else{
		fmt.Println("啥呀，垃圾东西")
		i++
		p.regret(i)
	}
}
func main() {
	p := &person{
		name: "student",
		age: 18,
		gender: "female",
	}
	var sentence string
	i:=0
	p.gu()
	p.sour()
	p.regret(i)
	fmt.Println("整句话来复读吧：")
	fmt.Scan(&sentence)
	p.repeat(sentence)
}