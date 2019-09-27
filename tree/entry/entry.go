package main

import "../../tree/node"

type MyNode struct {
	node *node.Node
}

// 扩展原有的 Node
func (myNode *MyNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := MyNode{myNode.node.Left}
	right := MyNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

// 方法或者是常量，首字母大写: public， 小写: private
func main() {
	var root node.Node
	root = node.Node{Value: 3}
	root.Left = &node.Node{}
	root.Right = &node.Node{5, nil, nil}
	root.Right.Left = new(node.Node)
	root.Print()

	//nodes := []treeNode{
	//	{value: 3},
	//	{},
	//	{6, nil, &root},
	//}
	//fmt.Println(nodes)

	root.Right.Left.SetValue(4) // 改不掉的，因为是值传递
	root.Right.Left.Print()
	root.Right.Left.SetValue2(4) // 传入指针，就直接将地址传递过去
	root.Right.Left.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue2(200)
	pRoot.Print()

	// nil 指针也可以调用方法
	var nilRoot *node.Node
	nilRoot.SetValue2(300)

	root.Traverse()

	myNode := MyNode{&root}
	myNode.postOrder()
}
