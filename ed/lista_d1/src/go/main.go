package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Node struct {
	Value int 
	next *Node
	prev *Node
}

type LList struct {
	root *Node
	size int
}
func NewLList() *LList{
	root := &Node{}
	root.next = root
	root.prev = root
	return &LList{root: root, size: 0}
} 
func (ll *LList) Size() int {
	return ll.size
}
func (ll *LList) String() string {
	text := "["
	node := ll.root.next
	for node != ll.root {
		text += fmt.Sprint(node.Value)
		if node.next != ll.root {
			text += ", "
		}
		node = node.next
	}
	text += "]"
	return text
}
func (ll *LList) PushFront(value int){
	newNode := &Node{Value: value}
	newNode.next = ll.root.next
	newNode.prev = ll.root	
	ll.root.next.prev = newNode
	ll.root.next = newNode
	ll.size++
}
func (ll *LList) PushBack(value int){
	newNode := &Node{Value: value}
	newNode.prev = ll.root.prev
	newNode.next = ll.root	
	ll.root.prev.next = newNode
	ll.root.prev = newNode
	ll.size++
}
func (ll *LList) PopFront(){
	if ll.size == 0{
		return
	}
	node := ll.root.next
	ll.root.next = node.next
	node.next.prev = ll.root
	ll.size--
}
func (ll *LList) PopBack(){
	if ll.size == 0{
		return
	}
	node := ll.root.prev
	ll.root.prev = node.prev
	node.prev.next = ll.root
	ll.size--
}
func (ll *LList) Clear(){
	ll.size = 0
	ll.root.next = ll.root
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
