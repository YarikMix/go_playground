package main

import "fmt"

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

	var stack []int
	registers := make(map[rune]int)

	for idx := 0; idx < len(program); idx++ {
		cmd := program[idx]

		switch cmd {
		case CmdPush:
			stack = append(stack, program[idx+1])
			idx++
		case CmdAdd:
			top := len(stack) - 1 // индекс верхнего элемента стека
			first := stack[top]   // получаем значение верхнего элемента
			stack = stack[:top]   // удаляем верхний элемент из стека
			top = len(stack) - 1  // индекс верхнего элемента стека
			second := stack[top]  // получаем значение верхнего элемента
			stack = stack[:top]   // удаляем верхний элемент из стека
			stack = append(stack, first+second)
		case CmdSub:
			top := len(stack) - 1 // индекс верхнего элемента стека
			y := stack[top]       // получаем значение верхнего элемента
			stack = stack[:top]   // удаляем верхний элемент из стека
			top = len(stack) - 1  // индекс верхнего элемента стека
			x := stack[top]       // получаем значение верхнего элемента
			stack = stack[:top]   // удаляем верхний элемент из стека
			stack = append(stack, x-y)
		case CmdMul:
			top := len(stack) - 1 // индекс верхнего элемента стека
			first := stack[top]   // получаем значение верхнего элемента
			stack = stack[:top]   // удаляем верхний элемент из стека
			top = len(stack) - 1  // индекс верхнего элемента стека
			second := stack[top]  // получаем значение верхнего элемента
			stack = stack[:top]   // удаляем верхний элемент из стека
			stack = append(stack, first*second)
		case CmdDiv:
			top := len(stack) - 1 // индекс верхнего элемента стека
			y := stack[top]       // получаем значение верхнего элемента
			stack = stack[:top]   // удаляем верхний элемент из стека
			top = len(stack) - 1  // индекс верхнего элемента стека
			x := stack[top]       // получаем значение верхнего элемента
			stack = stack[:top]   // удаляем верхний элемент из стека
			stack = append(stack, x/y)
		case CmdPrint:
			top := len(stack) - 1 // индекс верхнего элемента стека
			elem := stack[top]    // получаем значение верхнего элемента
			fmt.Println(elem)     // печатаем значение верхнего элемента
		case CmdSave:
			top := len(stack) - 1 // индекс верхнего элемента стека
			elem := stack[top]    // получаем значение верхнего элемента
			registers[rune(program[idx+1])] = elem
			idx++
		case CmdLoad:
			stack = append(stack, registers[rune(program[idx+1])])
			idx++
		case CmdPop:
			top := len(stack) - 1 // индекс верхнего элемента стека
			stack = stack[:top]   // удаляем верхний элемент из стека
		default:
			continue
		}
	}
}
