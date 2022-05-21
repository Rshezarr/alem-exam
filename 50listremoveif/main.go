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
	cur := l.Head
	for cur != nil {
		if cur.Data == data_ref {
			if prev == nil {
				l.Head = cur.Next
			} else {
				prev.Next = cur.Next
			}
		} else {
			prev = cur
		}
		cur = cur.Next
	}
}
