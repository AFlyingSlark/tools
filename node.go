package main

import "fmt"

type treeNode struct {
	value int
	left, right *treeNode
}

// 给 结构 定义 方法 !!!!!  也就是说 print setValue 是属于 treeNode的
// go 所有的参数都是传值的
// treeNode 类型的参数 abc，这个 abc 就称作接收者  相当于this self
func (abc treeNode) print()  {
// print 给 abc 接收的.
	fmt.Print(abc.value, " ")
}

func (aaa *treeNode) setValue(value int)  {
	// 即使是 指针  调用方式 还是 .
	// nil指针也可以调用方法
	if aaa == nil {
		fmt.Println("Setting value to nil" +
			"node. Ignored")
		return
	}
	aaa.value = value
}

// go 没有构造函数. 但可以用工厂函数代替
func createNode(value int) *treeNode {
	// 注意 返回的是局部变量的地址. 在其他语言会报错的.
	return &treeNode{value: value}
}


func main()  {
	var  root treeNode
	fmt.Println(root)

	root = treeNode{value: 3}
	fmt.Println(root)
	root.left = &treeNode{}
	fmt.Println(root)
	root.right = &treeNode{5, nil, nil}
	fmt.Println(root)
	// 一律使用 . 来访问成员
	root.right.left = new(treeNode)
	fmt.Println(root)

	nodes := []treeNode {
		{value:3},
		{},
		{6, nil, &root},	// 不能超过struct 定义的长度
		{},
		{5, &root, &treeNode{}},
	}
	fmt.Println(nodes)

	// 用到就放在堆上(有垃圾回收机制).   不用到就是在栈上(局部变量函数退出就销毁了)
	root.left.right = createNode(2)

	// abc 就是 root 相当于 this  如果 abc 写在 函数的括号里.那么就是 print(root)
	root.print()
	fmt.Println()

	root.right.left.setValue(4)
	// 还是 0 改不掉 因为是 值传递. 所以添加 * 把指针传进来
	root.right.left.print()

	// 不管是值, 指针 传递. 都可以用值(root) 来调用的
	root.print()
	root.setValue(100)

	pRoot := & root
	// 这个是值接收者 拷贝了一个pRoot
	pRoot.print()
	// 这个是指针接收者
	pRoot.setValue(200)
	pRoot.print()

	// 定义一个 nil类型的*treeNode tRoot
	var tRoot *treeNode
	tRoot.setValue(200)
	tRoot = &root
	tRoot.setValue(300)
	tRoot.print()
}
