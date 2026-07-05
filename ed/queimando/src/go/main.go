package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l int
	c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	for !stack.IsEmpty(){
		pos := stack.Pop()
		_ = pos
		if pos.l < 0 || pos.l >= len(grid) || pos.c < 0 || pos.c >= len(grid[0]) {
			continue
		}
	
		if grid[pos.l][pos.c] != '#' {
			continue
		}
	
		grid[pos.l][pos.c] = 'o'
		stack.Push(Pos{pos.l - 1, pos.c})
		stack.Push(Pos{pos.l + 1, pos.c})
		stack.Push(Pos{pos.l, pos.c - 1})
		stack.Push(Pos{pos.l, pos.c + 1})
	}

	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
