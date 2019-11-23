package main

import (
	"fmt"

	"goEasyDemo/tree"

	"golang.org/x/tools/container/intsets"
)

type customTree tree.NkTreeNode

// 相当于继承后重载函数的功能
func (cu *customTree) Print() {

	fmt.Println("custom!!!")
}

func testSparse() {

	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(100)
	s.Insert(10100)

	fmt.Println(s)
}

var root tree.NkTreeNode

func buildTree() {

	root = tree.NkTreeNode{Value: 3}
	root.Left = &tree.NkTreeNode{}
	root.Right = &tree.NkTreeNode{5, nil, nil}
	root.Right.Left = new(tree.NkTreeNode)
	root.Left.Right = tree.CreateNode(5)

	nodes := []tree.NkTreeNode{
		{Value: 3},
		{},
		{6, nil, &root},
	}

	fmt.Println(nodes)
	root.Print()
	root.SetVarValue(12) // 不改变值
	root.Print()
	root.SetPtrValue(12) // 改变值
	root.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetVarValue(15) // 不改变值
	pRoot.Print()
	pRoot.SetPtrValue(15) // 改变值
	pRoot.Print()
}

func main() {

	buildTree()

	// 中序遍历
	fmt.Println()
	root.Traverse()

	// 遍历一遍，并计数
	count := 0
	root.TraverseFunc(func(n *tree.NkTreeNode) {
		n.Print()
		count++
	})
	fmt.Println("count:", count)

	// 通过channel遍历，找到其中最大的节点
	c := root.TraverseWithChannel()
	maxNode := 0
	for node := range c { // 这里从channel中取出来node值
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}

	fmt.Println("Max node value: ", maxNode)

	// 继承自定义tree
	fmt.Println()
	cusTree := customTree{10, nil, nil}
	cusTree.Print()

	fmt.Println()
	testSparse()

}
