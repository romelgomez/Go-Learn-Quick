package node

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 工厂函数
func CreateNode(value int) *Node {
	return &Node{Value: value} // 局部变量也可返回出去，会自动分配到堆内存上
}

// 为结构定义方法, 也是值传递
func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node Node) SetValue(newValue int) {
	node.Value = newValue
}

func (node *Node) SetValue2(newValue int) {
	if node == nil {
		fmt.Println("setting nil node")
		return
	}
	node.Value = newValue
}

