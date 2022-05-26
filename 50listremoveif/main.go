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
	var prev *NodeL

	for l.Head != nil {
		if l.Head.Data == data_ref {
			if prev == nil {
				l.Head = l.Head.Next
			} else {
				prev.Next = l.Head.Next
			}
		} else {
			prev = l.Head
		}
		l.Head = l.Head.Next
	}
}
