package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
	l.size++
}

func (l *LList) Front() *Node {
	if l.size == 0 {
		return nil
	}
	return l.root.next
}

func (N *Node) Next() *Node {
	if N.next == N.root {
		return nil
	}
	return N.next
}

func addsorted(ll *LList, value int) {
	for node := ll.Front(); node != nil; node = node.Next(){
		if node.Value >= value {
			ll.insertBefore(node, value)
			return
		}
	}
	ll.PushBack(value)
}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
		
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}

func (l *LList) Reverse() {
	current := l.root
	
	for {
		current.next, current.prev = current.prev, current.next
		current = current.prev

		if current == l.root{
			break
		}
	}
}

func (l *LList) Merge(other *LList) *LList{
	merged := NewLList()
	nodeA := l.Front()
	nodeB := other.Front()

	for nodeA != nil && nodeB != nil {
		if nodeA.Value <= nodeB.Value {
			merged.PushBack(nodeA.Value)
			nodeA = nodeA.Next()
		} else {
			merged.PushBack(nodeB.Value)
			nodeB = nodeB.Next()
		}
	}

	for nodeA != nil {
		merged.PushBack(nodeA.Value)
		nodeA = nodeA.Next()
	}

	for nodeB != nil {
		merged.PushBack(nodeB.Value)
		nodeB = nodeB.Next()
	}

	return merged
}

func (l *LList) String() string {
	values := []string{}
	for node := l.Front(); node != nil; node = node.Next() {
		values = append(values, strconv.Itoa(node.Value))
	}

	return "[" + strings.Join(values, ", ") + "]"
}

func equals(a, b *LList) bool {
	if a.size != b.size {
		return false
	}
	nodeA := a.Front()
	nodeB := b.Front()
	for nodeA != nil && nodeB != nil {
		if nodeA.Value != nodeB.Value {
			return false
		}
		nodeA = nodeA.Next()
		nodeB = nodeB.Next()
	}
	return true
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla)
		case "reverse":
			lla := str2list(args[1])
			lla.Reverse()
			fmt.Println(lla)
		case "merge":
			lla := str2list(args[1])
			llb := str2list(args[2])
			merged := lla.Merge(llb)
			fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
