package main

import "fmt"

func main() {
	var n, e int
	fmt.Scan(&n, &e)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i + 1
	}

	p := e - 1

	for len(a) > 1 {
		fmt.Print("[ ")
		for i := 0; i < len(a); i++ {
			if i == p {
				fmt.Print(a[i], "> ")
			} else {

				fmt.Print(a[i], " ")
			}
		}
		fmt.Println("]")

		m := p + 1
		if m >= len(a) {
			m = 0
		}

		a = append(a[:m], a[m+1:]...)


		p = p + 1
		if p >= len(a) {
			p = 0
		}
	}

	fmt.Printf("[ %d> ]\n", a[0])
}