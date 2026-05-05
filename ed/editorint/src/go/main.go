package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	lines  *List[*List[rune]]
	line   *Node[*List[rune]]
	cursor *Node[rune]
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.cursor = e.line.Value.Insert(e.cursor, r)
	e.cursor = e.cursor.Next()
}

func (e *Editor) KeyLeft() {
	if e.cursor != e.line.Value.Front() { // Se o cursor não está no início da linha
		e.cursor = e.cursor.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.line != e.lines.Front() { // Se não está na primeira linha
		e.line = e.line.Prev()        // Move para a linha anterior
		e.cursor = e.line.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	newLine := NewList[rune]()

	current := e.cursor
	for current != e.line.Value.End() {
		next := current.Next()
		newLine.PushBack(current.Value)
		e.line.Value.Erase(current)
		current = next
	}

	insertedLine := e.lines.Insert(e.line.Next(), newLine)
	e.line = insertedLine
	e.cursor = e.line.Value.Front()
}

func (e *Editor) KeyRight() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.cursor.Next()
		return
	}

	if e.line.Next() != e.lines.End() {
		e.line = e.line.Next()
		e.cursor = e.line.Value.Front()
	}
}

func (e *Editor) KeyUp() {
	if e.line == e.lines.Front() {
		return
	}

	index := e.line.Value.IndexOf(e.cursor)
	e.line = e.line.Prev()
	e.cursor = moveToIndex(e.line.Value, index)
}

func (e *Editor) KeyDown() {
	if e.line == e.lines.Back() {
		return
	}

	index := e.line.Value.IndexOf(e.cursor)
	e.line = e.line.Next()
	e.cursor = moveToIndex(e.line.Value, index)
}

func (e *Editor) KeyBackspace() {
	if e.cursor != e.line.Value.Front() {
		e.line.Value.Erase(e.cursor.Prev())
		return
	}

	if e.line == e.lines.Front() {
		return
	}

	previous := e.line.Prev()
	e.cursor = appendLine(previous.Value, e.line.Value)

	e.lines.Erase(e.line)
	e.line = previous
}

func (e *Editor) KeyDelete() {
	if e.cursor != e.line.Value.End() {
		e.cursor = e.line.Value.Erase(e.cursor)
		return
	}

	if e.line.Next() == e.lines.End() {
		return
	}

	nextLine := e.line.Next()
	e.cursor = appendLine(e.line.Value, nextLine.Value)

	e.lines.Erase(nextLine)
}

func moveToIndex(list *List[rune], index int) *Node[rune] {
	node := list.Front()
	for i := 0; i < index && node != list.End(); i++ {
		node = node.Next()
	}
	return node
}

func appendLine(dest, src *List[rune]) *Node[rune] {
	var firstInserted *Node[rune]

	current := src.Front()
	for current != src.End() {
		inserted := dest.Insert(dest.End(), current.Value)
		if firstInserted == nil {
			firstInserted = inserted
		}
		current = current.Next()
	}

	return firstInserted
}

func main() {
	// Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.lines = NewList[*List[rune]]()
	e.lines.PushBack(NewList[rune]())
	e.line = e.lines.Front()
	e.cursor = e.line.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.lines.Front(); line != e.lines.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '⤶'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.cursor {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
