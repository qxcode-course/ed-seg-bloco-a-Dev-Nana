package main

import "fmt"

func main() {
	var h, p, f, d int32
	fmt.Scan(&h, &p, &f, &d)

	for {
		f = (f + 16) % 16

		if f == p {
			fmt.Println("N")
			break
		}

		if f == h {
			fmt.Println("S")
			break
		}

		if d == -1 {
			f--
		} else {
			f++
		}
	}
}