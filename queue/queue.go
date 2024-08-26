package main

import "fmt"

func main() {
	var ch int
	front := -1
	var queue []int
	for {
		fmt.Print("-----Choices-----\n0.exit\n1.enqueue\n2.dequeue\n3.sort\n4.search\nEnter your choice\n")
		fmt.Scan(&ch)
		if ch == 0 {
			break
		}
		switch ch {
		case 1:
			var item int
			fmt.Println("Enter data for queue")
			fmt.Scan(&item)
			front, queue = enqueue(item, front, queue)
			fmt.Println(queue)
			break
		case 2:
			if front == -1 {
				fmt.Println("queue is Empty....!")
			} else {
				queue, front = dequeue(front, queue)
				fmt.Println(queue)
			}
			break
		case 3:
			if front == -1 {
				fmt.Println("Queue is empty so can't perform sorting operation")
			} else {
				queue = sort(queue, front)
				fmt.Println(queue)
			}
			break
		case 4:
			if front == -1 {
				fmt.Println("Queue is empty so can't perform Searching operation")
			} else {
				var pos, s int
				fmt.Println("Enter element you want to search")
				fmt.Scan(&s)
				pos = search(queue, front, s)
				if pos != -1 {
					fmt.Println(s, "element founded at ", pos)
				} else {
					fmt.Println("Element not founded in Queue")
				}
			}
			break
		default:
			fmt.Println("Invalid choice")
			break
		}

	}

}
func enqueue(data, front int, queue []int) (int, []int) {
	queue = append(queue, data)
	front++
	return front, queue
}
func dequeue(front int, queue []int) ([]int, int) {
	var temp []int
	for i := 1; i <= front; i++ {
		temp = append(temp, queue[i])
	}
	front--
	return temp, front
}
func sort(queue []int, front int) []int {
	for i := 0; i <= front; i++ {
		for j := 0; j <= front; j++ {
			if queue[i] < queue[j] {
				ele := queue[i]
				queue[i] = queue[j]
				queue[j] = ele
			}
		}
	}
	return queue
}
func search(queue []int, front, s int) int {
	for i := 0; i <= front; i++ {
		if queue[i] == s {
			return i
		}
	}
	return -1
}
