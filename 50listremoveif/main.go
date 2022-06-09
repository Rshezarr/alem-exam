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
	if l.Head == nil {
		return
	}

	if l.Head.Data == data_ref {
		l.Head = l.Head.Next
		if l.Head == nil {
			return
		}
	}

	var prev *NodeL
	lh := l.Head
	for lh != nil {
		if lh.Data == data_ref {
			prev.Next = lh.Next
		} else {
			prev = lh
		}
		lh = lh.Next
	}
}
