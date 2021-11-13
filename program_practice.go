package main
// 用于练习 go 语言编程


import (
	"fmt";
	"math";
	"strings";
	"reflect"
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


// 获取单词个数 map
func get_word_count_map(s string) map[string]int {
	fmt.Println("进入获取单词个数 map")
	s_list := strings.Fields(s)
	fmt.Println("s_list", reflect.TypeOf(s_list))
	fmt.Println(s_list)

	word_count_map := make(map[string]int)
	for _, word := range s_list{
		fmt.Println(word)
		count, ok := word_count_map[word]
		if (ok) {
			word_count_map[word] = count + 1
		} else {
			word_count_map[word] = 1
		}
	}

	return word_count_map
}


// 斐波那契数
func get_geibonaqi_number() func () int  {
	a := 0
	b := 1
	closure := func () int {
		a, b = b, a + b
		return a
	}

	return closure
}


func main() {
	feibo_generate := get_geibonaqi_number()
	for i:=0; i<13; i++ {
		num := feibo_generate()
		fmt.Println(num)
	}
}
