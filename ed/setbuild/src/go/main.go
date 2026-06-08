package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
type Set struct {
	data []int
	size int
	capacity int
}
func NewSet(capacity int) *Set {
	return &Set{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (s *Set) Show () string{
	if s.size == 0 {
		return "[]"
	}

	text := "["
	for i := 0; i < s.size; i++ {
		if i > 0 {
			text += ", "
		}
		text += fmt.Sprint(s.data[i])
	}
	text += "]"
	return text
}
func (s *Set) Reserve (newCapacity int) {
	novo := make([]int, newCapacity)
	for i := 0; i < s.size; i++ {
		novo[i] = s.data[i]
	}
	s.data = novo
	s.capacity = newCapacity
}
func (s *Set) BinarSearch(value int) int {
	left := 0
	right := s.size - 1
	for left <= right {
		mid := (left + right) / 2
		if s.data[mid] == value {
			return mid
		}
		if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func (s *Set) Insert (value int){
	if s.capacity == 0 {
		s.Reserve(1)
	}
	if s.Contains(value) {
		return
	}
	if s.size == s.capacity {
		s.Reserve(s.capacity * 2)
	}
	pos := s.size
	for i := 0; i < s.size; i++ {
		if s.data[i] > value {
			pos = i
			break
		}
	}
	for i := s.size; i > pos; i--{
		s.data[i] = s.data[i-1]
	}
	s.data[pos] = value
	s.size++
}
func (s *Set) Contains (value int) bool{
	return s.BinarSearch(value) != -1
}
func (s *Set) Erase (value int) bool{
	index := s.BinarSearch(value)	
	if index == -1{
		fmt.Println("value not found")
		return false
	}
	for i := index; i < s.size-1; i++{
		s.data[i] = s.data[i+1]
	}
	s.size--
	return true
}
func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]
		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.Show())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			v.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			fmt.Println(v.Contains(value))
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}