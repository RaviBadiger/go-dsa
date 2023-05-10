package main
import "fmt"

type node struct{
	data int
	next *node
}

type linkedlist struct{
	head *node
	length int
}

func (l *linkedlist) prepend (n *node){
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedlist) printData(){
	toPrint := l.head
	for l.length != 0 {
		fmt.Println(toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println()
}

func (l *linkedlist) deleteWithValue(data int){
	if l.length == 0 {
		return
	}
	if l.head.data == data{
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != data{
		previousToDelete = previousToDelete.next
		if previousToDelete.next == nil {
			return
		}
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main(){
	myList := linkedlist{}
	node1 := &node{data: 48}
	node2 := &node{data: 38}
	node3 := &node{data: 28}
	node4 := &node{data: 18}
	node5 := &node{data: 31}
	node6 := &node{data: 24}
	node7 := &node{data: 13}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)
	myList.printData()
	myList.deleteWithValue(31)
	myList.deleteWithValue(310)
	myList.deleteWithValue(13)
	myList.printData()
	emptyList := linkedlist{}
	emptyList.deleteWithValue(10)


}