package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	var nl, nc int

	scanner.Scan()
	fmt.Sscan(scanner.Text(), &nl, &nc)

	grid := make([][]rune, nl)
	visitado := make([][]bool, nl)

	for i := range visitado {
		visitado[i] = make([]bool, nc)

		grid[i] = make([]rune, nc)
	}
	
	var inicio, fim Pos
	
	for i := 0; i < nl; i++ {
		scanner.Scan()
		linha := scanner.Text()

		for j := 0; j < nc && j < len(linha); j++{
			grid[i][j] = rune(linha[j])

			if grid[i][j] == 'I' {
				inicio = Pos{i, j}
			} else if grid[i][j] == 'F' {
				fim = Pos{i, j}
			}
		}
	}
	caminho := NewStack[Pos]()
	becos := NewStack[Pos]()

	caminho.Push(inicio)

	dl := []int{-1, 0, 1, 0}
	dc := []int{0, 1, 0, -1}

	for !caminho.IsEmpty() {
		atual := caminho.Top()
		visitado[atual.l][atual.c] = true

		if atual == fim {
			break
		}

		achou := false
		
		for i := 0; i < 4; i++{
			nl2 := atual.l + dl[i]
			nc2 := atual.c + dc[i]

			if nl2 < 0 || nl2 >= nl || nc2 < 0 || nc2 >= nc {
				continue
			}

			if grid[nl2][nc2] != '#' && !visitado[nl2][nc2]{
				caminho.Push(Pos{nl2, nc2})
				achou = true
				break
			}
		}

		if !achou {
			beco := caminho.Pop()

			if grid[beco.l][beco.c] != 'I' && grid[beco.l][beco.c] != 'F' {
				grid[beco.l][beco.c] = 'x'
			}

			becos.Push(beco)
		}
	}
	
	for !becos.IsEmpty() {
		p := becos.Pop()
		if grid[p.l][p.c] == 'x' {
			grid[p.l][p.c] = ' '
		}
	}

	for !caminho.IsEmpty () {
		p := caminho.Pop()
		grid[p.l][p.c] = '.'
	}

	for _, linha := range grid {
		fmt.Println(string(linha))
	}
}