package node

// 一个 struct 不同的方法可以分散在不同的文件里
func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
