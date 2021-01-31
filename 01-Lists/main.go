package main
/*
	A list is a sequence of elements. Each element can be connected to
	another with a link in a forward or backward direction. You can add
	or remove elements easily.
*/
import (
	"container/list"
	"fmt"
)

func main() {
	//create new list and put some items
	l := list.New() //creates list

	item4 := l.PushBack(4) //put item on back --> [4]
	item1 := l.PushFront(1) //put item on front --> [1, 4]

	l.InsertBefore(3, item4) //put item before item4:=4 --> [1, 3, 4]
	l.InsertAfter(2, item1) //put item after item1:=5 --> [1, 2, 3, 4]

	for item := l.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value)
	}
}
