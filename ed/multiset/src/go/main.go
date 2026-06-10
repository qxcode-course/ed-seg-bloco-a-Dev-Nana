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

func NewMultiSet(capacity int) *MultiSet{
	return &MultiSet{
		data: make([]int, 0, capacity),
		size: 0}
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

	for i := ms.size; i > index; i--{
		ms.data[i] = ms.data[i-1]
	}

	ms.data[index] = value
	ms.size++
}

func (ms *MultiSet) Insert(value int){
	_, index := ms.Search(value)
	ms.insert(value, index)
}

func (ms *MultiSet) Contains(value int) bool {
	found, _ := ms.Search(value)
	return found
}

func (ms *MultiSet) erase(index int) {
	for i := index; i < ms.size - 1; i++ {
		ms.data[i] = ms.data[i+1]
	}

	ms.size--
}

func (ms *MultiSet) Erase(value int) error {
	found, index := ms.Search(value)
	if !found {
		return fmt.Errorf("value not found")
	}

	ms.erase(index)
	return nil
}

func (ms *MultiSet) Count(value int) int {
	found, index := ms.Search(value)

	if !found {
		return 0
	}

	count := 0 

	for i := index; i < ms.size && ms.data[i] == value ; i++{
		count++
	}

	return count
}

func (ms *MultiSet) String() string {
	return "[" + Join(ms.data[:ms.size], ", ") + "]"
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

func (ms *MultiSet) Unique() int {
	if ms.size == 0{
		return 0
	}

	count := 1

	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1]{
			count ++
		}
	}

	return count
}

func (ms *MultiSet) Clear() {
	ms.size = 0
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
			value, _ := strconv.Atoi(args[1])
			err := ms.Erase(value)
			if err != nil{
				fmt.Println(err)
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
