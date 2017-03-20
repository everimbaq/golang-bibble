package LinkedList

import "fmt"

type Node struct {
	next *Node
	Data interface{}
}

type ListData interface {}

type MyLinkedList *Node

func (mynode *Node)FindNth(pos int) ListData{

	if mynode == nil{
		return nil
	}

	if pos <=1{
		return mynode.Data
	}
	return mynode.next.FindNth(pos -1)

}

func (node *Node)Insert(pos int, data ListData) error {
	curNode := node
	for tmp :=1;;tmp++{

		if tmp == pos{
			newNode := Node{curNode.next, data}
			curNode.next = &newNode
			return nil
		}else if tmp < pos{

			if curNode.next == nil && tmp +1 == pos{
				newNode := Node{curNode.next, data}
				curNode.next = &newNode
				return nil
			}else if curNode.next != nil{
				curNode = curNode.next
			}else {
				return fmt.Errorf("list is not long enough")
			}
		}else{
			return fmt.Errorf("list is not long enough")
		}

	}
}

func(node *Node)Delete(pos int) (interface{}, error ){
	curNode := node
	var tmp int = 1
	for curNode != nil && tmp < pos-1{
		curNode = curNode.next
		tmp++
	}

	if curNode == nil || tmp > pos-1{
		return nil, fmt.Errorf("list is not long enough")
	}

	tmpNode := curNode.next
	if tmpNode == nil {
		return nil, fmt.Errorf("list is not long enough")

	}else {
		curNode.next = tmpNode.next
		return tmpNode.Data ,nil
	}

}