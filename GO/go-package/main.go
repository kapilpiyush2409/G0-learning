package main

import (
	"fmt"
	"project/go/dataStructure"
)


	func main() {
      // Creating a new binary tree
      
      tree := datastructure.CreateBST()
  
      // Inserting elements into the tree
      tree.Insert(5)
      tree.Insert(3)
      tree.Insert(7)
      tree.Insert(1)
      tree.Insert(4)
      tree.Insert(6)
      tree.Insert(8)
      // Printing the tree using in-order traversal
      fmt.Println("In-order traversal of the binary tree:")
      tree.InOrderTraversal()
      fmt.Println()
      tree.PreOrderTraversal()
      fmt.Println()
      fmt.Println(tree.Root.Find(7))
      fmt.Println(tree.Root.Find(77))
      // tree.InOrderTraversal(tree.Root)
      // tree.Root=tree.Root.Delete(7)
      // fmt.Println()
      // tree.InOrderTraversal(tree.Root)
      
  }

