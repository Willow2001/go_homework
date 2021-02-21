package main

import "fmt"

// Author 用户
type Author struct {
	Name string //名字
	VIP bool //是否是高贵的带会员
	Icon string //头像
	Signature string //签名
	Focus int //关注人数
}

//VideoInformation 视频详情
type VideoInformation struct {
	AuthorName string //作者名
	VideoName string //视频名
	LikeNum int //点赞数
	CoinNum int//投币数
	CollectNum int//收藏数
}

// Like 点赞
func (a *Author) Like(v *VideoInformation) {
	v.LikeNum++
	fmt.Println(a.Name,"点赞了该视频")
}
// Coin 投币
func (a *Author) Coin(v *VideoInformation) {
	v.CoinNum++
	fmt.Println(a.Name,"投币了该视频")
}
// Collect 收藏
func (a *Author) Collect(v *VideoInformation) {
	v.CollectNum++
	fmt.Println(a.Name,"收藏了该视频")
}
//Teak 一键三连
func (a *Author) Teak(v *VideoInformation) {
	v.LikeNum++
	v.CoinNum++
	v.CollectNum++
	fmt.Println(a.Name,"一键三连了该视频")
}

//发布视频
func PostVideo(Aname, Vname string) VideoInformation {
	fmt.Println("著名鸽子王",Aname,"发布了视频")
	V :=VideoInformation{
		AuthorName :Aname,
		VideoName :Vname,
		LikeNum :0,
		CoinNum :0,
		CollectNum :0,
	}
	return V
}
func main() {
	V := PostVideo("罗翔", "张三说刑法")
	p := Author{
		Name :"student",
		VIP :true,
		Icon :"鸽子",
		Signature :"学不动了",
		Focus :1000,
	}

	p.Like(&V)
	p.Coin(&V)
	p.Collect(&V)
	p.Teak(&V)
	fmt.Println(V)
}

