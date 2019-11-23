package tree

import (
	"fmt"
)

type NkTreeNode struct {
	Value       int
	Left, Right *NkTreeNode
}

func (node NkTreeNode) Print() {
	fmt.Printf("%d \n", node.Value)
}

func CreateNode(value int) *NkTreeNode {
	return &NkTreeNode{Value: value}
}

// 指针传递
func (node *NkTreeNode) SetPtrValue(va int) {
	node.Value = va
}

// 值传递，不改变原变量
func (node NkTreeNode) SetVarValue(va int) {
	node.Value = va
}

// 中序遍历
func (node *NkTreeNode) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

// 中序遍历树，对每个节点执行函数 f
func (node *NkTreeNode) TraverseFunc(f func(node *NkTreeNode)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *NkTreeNode) TraverseWithChannel() chan *NkTreeNode {
	out := make(chan *NkTreeNode)
	go func() {
		// 中序遍历树，对每个节点执行函数
		node.TraverseFunc(func(node *NkTreeNode) {
			// 这里把node塞到 channel中，外面会同步取出来，取出来之后进行其他操作
			out <- node
		})

		// 遍历完之后，关闭channel，外部的 for range channel 循环能够结束
		close(out)
	}()
	return out
}
