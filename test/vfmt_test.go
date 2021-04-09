package test

import (
	"github.com/vua/vfmt"
	"testing"
)

type user struct {
	id   int
	name string
	age  int
}

func Test(t *testing.T) {

	//无样式
	vfmt.Println("Console Output")
	vfmt.Println("@[Console]@[ Output::]")
	//无内容
	vfmt.Println("@[::red]")
	//有内容有样式
	vfmt.Println("@[Console Output::red]")
	//struct 输出
	u:=user{1,"vua",18}
	vfmt.Printf("@[%T,%v::#00ff00|bg#ff0000|bold]\n",u,u)
	//自定义样式
	//创建样式：超链接
	vfmt.RegisterStyle("url", "blue|underline")
	vfmt.Println("@[https://www.github.com::url]")
}
