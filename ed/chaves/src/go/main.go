package main

import "fmt"

func main() {
	fila := NewQueue[string]()
	for letra := 'A'; letra <= 'P'; letra++{
		fila.Enqueue(string(letra))
	}

	for i := 0; i < 15; i++ {
		time1 := fila.Dequeue()
		time2 := fila.Dequeue()
		var gols1, gols2 int
		fmt.Scan(&gols1, &gols2)

		if gols1 > gols2 {
			fila.Enqueue(time1)
		} else {
			fila.Enqueue(time2)
		}
	}

	fmt.Println(fila.Dequeue())
}
