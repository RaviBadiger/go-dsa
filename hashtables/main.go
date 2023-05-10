package main

import (
	"fmt"
)

const ArraySize = 7

// Hashtable structure
type Hashtable struct {
	array [ArraySize]*bucket

}
//bucket structure
type bucket struct{
	head *bucketNode
}
//bucketNode structure
type bucketNode struct{
	key string
	next *bucketNode
}
//Insert
func (h *Hashtable) Insert(key string){
	index := hash(key)
	h.array[index].insert(key)
}

//Search
func (h *Hashtable) Search(key string) bool{
	index := hash(key)
	return h.array[index].search(key)

	
}
//Delete
func (h *Hashtable) Delete(key string){
	index := hash(key)
	h.array[index].delete(key)
}

//insert
func (b *bucket) insert(k string){
	if !b.search(k){
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	}else {
		fmt.Println("Already exists")
	}
}
//search
func (b *bucket) search(k string) bool{
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}
//delete
func (b *bucket) delete(k string){
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next!= nil{
		if previousNode.next.key == k {
			previousNode = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}


//hash
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}
//Init
func Init() *Hashtable{
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
func main(){
 hashTable := Init()
 list := []string{
	"ERIC",
	"KENNY",
"KYLE",
"STAN",
"RANDY",
"BUTTERS",
"TOKEN" }
for _, v := range list {
	hashTable.Insert(v)
}

}