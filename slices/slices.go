package main

import "fmt"

func main() {
	var a []int
	var len, data, ch int
	fmt.Println("Enter how many elements you want to enter ")
	fmt.Scan(&len)
	for i := 0; i < len; i++ {
		fmt.Println("Enter data into slice a[", i, "]")
		fmt.Scan(&data)
		a = append(a, data)
	}
	fmt.Printf("0.exit \n 1.add \n 2. subtract \n 3. multiplaction \n 4. division \n 5. modulues \n 6.sort \n 7.search")
	fmt.Println("Enter Your choice")
	fmt.Scan(&ch)
	switch ch {
	case 0:
		break
	case 1:
		add(a)
		break
	case 2:
		sub(a)
		break
	case 3:
		mul(a)
		break
	case 4:
		div(a)
		break
	case 5:
		mod(a)
		break
	case 6:
		sorting(a, len)
		break
	case 7:
		searching(a, len)
		break
	default:
		fmt.Println("Invalid choice please enter index number")
	}
	fmt.Println("Enter choice")
}
func add(data []int) {
	add := 0
	for i := range data {
		add += data[i]
	}
	fmt.Println(" after adding all slice elements is", add)
}
func sub(data []int) {
	sub := 0
	for i := range data {
		if i == 0 {
			sub = data[0]
		} else {
			sub -= data[i]
		}
	}
	fmt.Println("after subtracting all slice elements is", sub)
}
func mul(data []int) {
	mul := 0
	for i := range data {
		if i == 0 {
			mul = data[0]
		} else {
			mul *= data[i]
		}
	}
	fmt.Println("after Multiply all slice elements is", mul)
}
func div(data []int) {
	div := 0
	for i := range data {
		if i == 0 {
			div = data[0]
		} else {
			div /= data[i]
		}
	}
	fmt.Println("after divides all slice elements is", div)
}
func mod(data []int) {
	mod := 0
	for i := range data {
		if i == 0 {
			mod = data[0]
		} else {
			mod %= data[i]
		}
	}
	fmt.Println("after subtracting all slice elements is", mod)
}
func sorting(data []int, len int) {
	var temp int
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if data[i] > data[j] {
				temp = data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}
	fmt.Println("descending sorted slice is", data)
	for i := 0; i < len; i++ {
		for j := 0; j < len; j++ {
			if data[i] < data[j] {
				temp = data[i]
				data[i] = data[j]
				data[j] = temp
			}
		}
	}
	fmt.Println("ascending sorted slice is", data)
}
func searching(data []int, len int) {
	var s, count int
	fmt.Println("Enter which element You want to search ")
	fmt.Scan(&s)
	for i := 0; i < len; i++ {
		if s == data[i] {
			fmt.Println(data[i], "Element is Founded at", count, "Position")
			break
		} else if i == len-1 {
			if s != data[i] {
				fmt.Println("Element is Not founded at any position")
			}
		}
		count++
	}
}
