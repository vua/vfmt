package test

import (
	"fmt"
	"github.com/vua/vfmt"
	"testing"
)

type user struct {
	id   int
	name string
	age  int
}

func Test(t *testing.T) {

	//有内容用样式
	output := "Console Output"
	style := "red"
	vfmt.Printf("[vfmt] @[%s::%s]\n", output, style)
	vfmt.Println("[vfmt] @[Console Output::red]")

	//有内容无样式 (四者等价)
	vfmt.Println("[vfmt] @[Console Output::]")
	vfmt.Println("[vfmt] @[Console Output]")
	vfmt.Println("[vfmt] Console Output")
	fmt.Println("[vfmt] Console Output")

	//有样式无内容 (五者等价)
	vfmt.Println("[vfmt] @[::red]")
	vfmt.Println("[vfmt] @[::]")
	vfmt.Println("[vfmt] @[]")
	vfmt.Println("[vfmt]")
	fmt.Println("[vfmt]")

	//结构体输出
	/*
		type user struct {
			  id   int
			  name string
			  age  int
		}
	*/
	u := user{1, "vua", 18}
	vfmt.Printf("[vfmt] @[%T,%v::#00ff00|bg#ff0000|bold]\n", u, u)

	//叠加样式
	vfmt.Println("[vfmt] @[Console Output::green|bgRed|bold]")

	//16进制颜色
	vfmt.Println("[vfmt] @[Console Output::#00ff00|bg#ff0000|bold]")

	//自定义样式
	//创建样式：超链接
	vfmt.RegisterStyle("url", "blue|underline")
	vfmt.Println("[vfmt] @[https://www.github.com::url]")

	//嵌套测试 (惰性匹配)
	vfmt.Println("[vfmt] @[@[Console Output::yellow]::red]")

	//并行测试 (惰性匹配)
	vfmt.Println("[vfmt] @[Console Output::yellow]@[Console Output::#cc0ffe]")
}
