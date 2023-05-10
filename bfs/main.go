package main

import "fmt"

//Node
type Node struct{
	Key int
	Left *Node
	Right *Node

}
//Insert
func (n *Node) Insert(k int){
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		}else{
			n.Right.Insert(k)
		}
	}else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		}else{
			n.Left.Insert(k)
		}
	}
}


//Search
func (n *Node)  Search(k int) bool {

	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	}else if n.Key > k {
		return n.Left.Search(k)
	}
	return true
}


func main(){
	tree := &Node{Key: 100}
	tree.Insert(200)
	tree.Insert(300)
	//fmt.Print(tree)
	tree.Insert(46)
	tree.Insert(76)
	tree.Insert(86)
	tree.Insert(96)
	fmt.Print(tree.Search(3001))

	
}

