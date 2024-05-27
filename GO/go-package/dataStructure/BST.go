package datastructure

import "fmt"



type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

type BinaryTree struct {
	Root *BinaryNode
}



func CreateBST() *BinaryTree{
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(data int){
	fmt.Println(&data)
	if &data == nil {
		panic("cannot be nil")
	}

	newNode := BinaryNode{Value: data,Left: nil,Right: nil}

	if bt.Root == nil{
		bt.Root = &newNode
	}else{
		bt.Root.insertInNode(newNode)
	}
}


func (node *BinaryNode) insertInNode(newNode BinaryNode){
	if node.Value > newNode.Value{
		if node.Left != nil{
			node.Left.insertInNode(newNode)
		}else{
			node.Left = &newNode
		}
		
	}else if node.Value < newNode.Value{
		if node.Right != nil{
			node.Right.insertInNode(newNode)
		}else{
			node.Right = &newNode
		}
		
	} 
}



func (bt BinaryTree) InOrderTraversal(){
	bt.inOrderTraversal(bt.Root)
}

func (bt BinaryTree) PreOrderTraversal(){
	bt.preOrderTraversal(bt.Root)
}


func (bt *BinaryTree) inOrderTraversal(node *BinaryNode) {
    if node != nil {
        bt.inOrderTraversal(node.Left)
        fmt.Printf("%d ", node.Value)
        bt.inOrderTraversal(node.Right)
    }
}

func (bt BinaryTree) preOrderTraversal(node *BinaryNode){
	if node != nil{
		
        fmt.Printf("%d ", node.Value)
		bt.preOrderTraversal(node.Left)
        bt.preOrderTraversal(node.Right)
	}
	
}


func (node *BinaryNode) Find(data int) bool{

	if node.Value == data{
		return true
	}else if node.Value < data && node.Left != nil{
		node.Left.Find(data)
	}else if node.Value > data && node.Right != nil{
		node.Right.Find(data)
	}
	return false

}