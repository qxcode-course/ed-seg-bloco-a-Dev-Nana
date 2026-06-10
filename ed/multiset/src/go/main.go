package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type MultiSet struct {
	data []int
	size int
}

func NewMultiSet(value int) *MultiSet{
	return &MultiSet{data: []int{value}, size: 0}
}

func (ms *MultiSet) Search (value int) (bool, int) {
	low := 0
	high := ms.size

	for low < high {
		mid := (low + high) / 2
		
		if ms.data[mid] < value{
			low = mid + 1
		} else {
			high =mid
		}
	}

	if low < ms.size && ms.data[low] == value {
		return true, low
	}

	return false, low
}

func (ms *MultiSet) insert(value int, index int) {
	ms.data = append(ms.data, 0)

	for i := ms.size; i > 0; i--{
		ms.data[i] = ms.data[i-1]
	}

	ms.data[0] = value
	ms.size++
}

func (ms *MultiSet) Insert(value int){
	_, index := ms.Search(value)
	ms.insert(value, index)
}

func (ms *MultiSet) String() string {
	return "[" + Join(ms.data, ", ") + "]"
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms)
		case "erase":
			// value, _ := strconv.Atoi(args[1])
		case "contains":
			// value, _ := strconv.Atoi(args[1])
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
