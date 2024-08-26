package main

import (
	"fmt"
)

func push(data, top int, stack []int) (int, []int) {
	stack = append(stack, data)
	top++
	return top, stack
}
func pop(top int, stack []int) ([]int, int) {
	var temp []int
	for i := 0; i < top; i++ {
		temp = append(temp, stack[i])
	}
	top--
	return temp, top
}
func sort(stack []int, top int) []int {
	for i := 0; i <= top; i++ {
		for j := 0; j <= top; j++ {
			if stack[i] < stack[j] {
				temp := stack[i]
				stack[i] = stack[j]
				stack[j] = temp
			}

		}
	}
	return stack
}
func search(stack []int, top, search int) int {
	for i := 0; i <= top; i++ {
		if stack[i] == search {
			return i
		}
	}
	return -1
}
func main() {
	var ch int
	top := -1
	var stack []int
	for {
		fmt.Print("-----Choices-----\n0.exit\n1.Push\n2.pop\n3.sort\n4.search\nEnter your choice\n")
		fmt.Scan(&ch)
		if ch == 0 {
			break
		}
		switch ch {
		case 1:
			var item int
			fmt.Println("Enter data for stack")
			fmt.Scan(&item)
			top, stack = push(item, top, stack)
			fmt.Println(stack)
			break
		case 2:
			if top == -1 {
				fmt.Println("Stack is Empty....!")
			} else {
				stack, top = pop(top, stack)
				fmt.Println(stack)
			}
			break
		case 3:
			if top == -1 {
				fmt.Println("Stack is empty so can't perform sorting operation")
			} else {
				stack = sort(stack, top)
				fmt.Println(stack)
			}
			break
		case 4:
			if top == -1 {
				fmt.Println("Stack is empty so can't perform Searching operation")
			} else {
				var pos, s int
				fmt.Println("Enter element you want to search")
				fmt.Scan(&s)
				pos = search(stack, top, s)
				if pos != -1 {
					fmt.Println(s, "element founded at ", pos)
				} else {
					fmt.Println("Element not founded in stack")
				}
			}
			break
		default:
			fmt.Println("Invalid choice")
			break
		}

	}

}
