package main

import "fmt"

type treeNode struct {
	varlue int
	left, right *treeNode
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
}
