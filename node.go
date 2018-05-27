package main

import "fmt"

type treeNode struct {
	varlue int
	left, right *treeNode
}
// go 没有构造函数. 但可以用工厂函数代替

func createNode(value int) *treeNode {
	// 注意 返回的是局部变量的地址. 在其他语言会报错的.
	return &treeNode{value:value}
}


func main()  {
	var  root treeNode
	fmt.Println(root)

	root = treeNode{varlue: 3}
	fmt.Println(root)
	root.left = &treeNode{}
	fmt.Println(root)
	root.right = &treeNode{5, nil, nil}
	fmt.Println(root)
	// 一律使用 . 来访问成员
	root.right.left = new(treeNode)
	fmt.Println(root)

	nodes := []treeNode {
		{varlue:3},
		{},
		{6, nil, &root},	// 不能超过struct 定义的长度
		{},
		{5, &root, &treeNode{}},
	}
	fmt.Println(nodes)

	// 用到就放在堆上(有垃圾回收机制).   不用到就是在栈上(局部变量函数退出就销毁了)
	root.left.right = createNode(2)

}
