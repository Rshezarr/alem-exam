package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	var tmpList List
	for node := l.Head; node != nil; node = node.Next {
		if node.Data != data_ref {
			//	ListPushBack(&tmpList, node.Data)
		}
	}
	l.Head = tmpList.Head
}
