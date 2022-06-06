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
	lh := l.Head
	for lh != nil {
		if lh.Data == data_ref {
			if prev == nil {
				lh = lh.Next
			} else {
				prev.Next = lh.Next
			}
		} else {
			prev = lh
		}
		lh = lh.Next
	}
}
