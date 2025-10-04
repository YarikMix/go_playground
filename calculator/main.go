package main

import (
	"calculator/stack"
	"fmt"
)

const (
	CmdAdd   = iota // сложить два числа в стеке (top-1) = (top-1) + top
	CmdSub          // вычесть (top-1) = (top-1) - top
	CmdMul          // умножить (top-1) = (top-1) * top
	CmdDiv          // разделить (top-1) = (top-1) / top
	CmdPush         // поместить следующее число в стек
	CmdPop          // убрать число из стека
	CmdPrint        // печать верхнего элемента стека
	CmdSave         // сохранить число из стека в ячейку
	CmdLoad         // переместить число из ячейки в стек
)

func main() {
	program := []int{CmdPush, 33, CmdPush, 44, CmdAdd, CmdPush, 567, CmdSub, CmdPush,
		-13, CmdMul, CmdPush, 5, CmdDiv, CmdPush, 45, CmdPush, 21, CmdAdd, CmdMul,
		CmdPrint, CmdSave, 'А', CmdPop, CmdPush, 3, CmdPush, 9, CmdPush, 7,
		CmdSub, CmdMul, CmdLoad, 'А', CmdMul, CmdPrint, CmdSave, 'Б',
		CmdLoad, 'А', CmdPush, 10230, CmdLoad, 'Б', CmdSub, CmdSub,
		CmdPush, 1000, CmdDiv, CmdPrint}

	var s stack.Stack
	registers := make(map[rune]int)

	for idx := 0; idx < len(program); idx++ {
		cmd := program[idx]

		switch cmd {
		case CmdPush:
			s = s.Push(program[idx+1])
			idx++
		case CmdAdd:
			var x, y int
			s, x = s.Pop()
			s, y = s.Pop()
			s = s.Push(x + y)
		case CmdSub:
			var x, y int
			s, y = s.Pop()
			s, x = s.Pop()
			s = s.Push(x - y)
		case CmdMul:
			var x, y int
			s, x = s.Pop()
			s, y = s.Pop()
			s = s.Push(x * y)
		case CmdDiv:
			var x, y int
			s, y = s.Pop()
			s, x = s.Pop()
			s = s.Push(x / y)
		case CmdPrint:
			elem := s.Pick()
			fmt.Println(elem) // печатаем значение верхнего элемента
		case CmdSave:
			elem := s.Pick()
			registers[rune(program[idx+1])] = elem
			idx++
		case CmdLoad:
			s = s.Push(registers[rune(program[idx+1])])
			idx++
		case CmdPop:
			s, _ = s.Pop()
		default:
			continue
		}
	}
}
