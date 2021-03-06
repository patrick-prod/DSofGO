package TREE

import (
	"fmt"
	"go_DataStruct/QUEUE"
	"go_DataStruct/STACK"
	"math"
)

type BinaryTree struct{
	Root	*BinaryNode
	Level	int
	i 		int
}

func NewEmptyBinaryTree() *BinaryTree {
	return &BinaryTree{nil,0,0}
}

func NewBinaryTree(preList []interface{}) *BinaryTree {
	bt := NewEmptyBinaryTree()
	bt.Root = bt.create(preList)
	bt.i = 0
	for _,item := range preList {
		if item != nil {
			bt.i++
		}
	}
	bt.Level = int(math.Floor(math.Log(float64(bt.i))+float64(1)))
	bt.i = 0
	return bt
}


func (bt *BinaryTree)create(values []interface{}) *BinaryNode {
	p := new(BinaryNode)
	if bt.i < len(values) {
		ele := values[bt.i]
		bt.i ++
		if ele != nil {
			p = &BinaryNode{Data:ele}
			p.Left = bt.create(values)
			p.Right = bt.create(values)
		}
	}
	return p
}

func (bt *BinaryTree) IsEmpty() bool {
	return bt.Root == nil
}

func (bt *BinaryTree) Insert(x interface{},parent *BinaryNode,isleft bool) *BinaryNode {
	if x == nil {
		return nil
	}

	if isleft {
		parent.Left = &BinaryNode{Data:x,Left:parent.Left,Right:nil}
		return parent.Left
	}else{
		parent.Right = &BinaryNode{Data:x,Left:nil,Right:parent.Right}
		return parent.Right
	}
}

func (bt *BinaryTree) Remove(parent *BinaryNode,isleft bool) {
	if isleft {
			parent.Left = nil
	}else{
			parent.Right = nil
	}
}

func (bt *BinaryTree) Clear() {
	bt.Root = nil
}

/** 递归遍历 */
func (bt *BinaryTree) PreOrder(x *BinaryNode) {
	for x != nil {
		fmt.Print(x.Data, " ")
		bt.PreOrder(x.Left)
		bt.PreOrder(x.Right)
	}
}

func (bt *BinaryTree) InOrder(x *BinaryNode) {
	for x != nil {
		bt.InOrder(x.Left)
		fmt.Println(x.Data, " ")
		bt.InOrder(x.Right)
	}
}

func (bt *BinaryTree) PostOrder(x *BinaryNode) {
	for x != nil {
		bt.InOrder(x.Left)
		bt.InOrder(x.Right)
		fmt.Println(x.Data, " ")
	}
}

/** 获取节点数 */
func (bt *BinaryTree) GetNodeNum(x *BinaryNode) int {
	if x == nil {
		return 0
	}else{
		a := bt.GetNodeNum(x.Left)
		b := bt.GetNodeNum(x.Right)
		return a + b + 1
	}
}

/** 中根非递归遍历 */
func (bt *BinaryTree) InOrderTraverse() {
	if bt.Root == nil {
		return
	}
	b  := bt.Root
	stack := STACK.NewSeqStack(bt.GetNodeNum(bt.Root))
	for b != nil || !stack.IsEmpty() {
		if b != nil {
			stack.Push(b)
			b = b.Left
		}else{
			b = stack.Pop().(*BinaryNode)
			fmt.Println(b.Data)
			b = b.Right
		}
	}
}

/** 层级遍历 */
func (bt *BinaryTree) LevelTraverse() {
	x := bt.Root
	queue := QUEUE.NewLinkedQueue()
	for x != nil {
		fmt.Print(x.Data," ")
		if x.Left  != nil { queue.Add(x.Left) }
		if x.Right != nil { queue.Add(x.Right) }
		x = queue.Poll().(*BinaryNode)
	}
}