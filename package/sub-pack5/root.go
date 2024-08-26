package Root

import "fmt"

func Root(num int) {

	fmt.Println(num, "is Root number of", root(num))
}
func root(data int) (root int) {
	root = data * data
	return
}
