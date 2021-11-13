package main
// 用于练习 go 语言编程


import (
	"fmt";
	"math"
)

func print_practice()  {
	fmt.Println("hello")
}


// 平方根函数
func get_square_root(x float64) float64 {
	left := 0.0
	right := x

	for {
		middle := (left + right) / 2
		fmt.Println(middle)
		if (math.Abs(math.Pow(middle, 2) - x) < 0.1) {
			break
		}

		if (middle*middle > x) {
			right = middle
		} else {
			left = middle
		}

	}

	return right
}

func main() {
	get_square_root(11024)
}