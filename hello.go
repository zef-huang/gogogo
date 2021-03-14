package main

import (
	"fmt"
	"strings"
	"net/http"
	// "math/rand"
	// "strconv"
	// "math/rand"
	// xixi "./utils"
	// "time"
// 	"sort"
)

type fun_nickname func(int) int

func peach_count(count int) int {
	if count==1 {
		return 1
	} else {
		return (peach_count(count-1) + 1) * 2
	}
}

func get_func_param(get_peach_count fun_nickname, day int) {
	count := get_peach_count(day)
	fmt.Println("10天前桃子数: ", count)
}

func test_func_return_params(num1 int, num2 float32)(sum1 int, sum2 float32){
	sum1 = int(num1) + int(num2)
	sum2 = float32(num1) + float32(num2)
	return
}

func swap_value(n1 *int, n2 *int) {
	tmp := *n1
	*n1 = *n2
	*n2 = tmp
}

func Adder (suffix string) (func (string) string) {
	return func (pic string) string {
		if strings.HasSuffix(pic, suffix){
			return pic
		} else {
			return pic + ".jpg"
		}
	}
}

// 长度也是类型的一部分
func change_arr(arr *[34]int) {
	(*arr)[33] = 100
}

func feibo(n int) []int {
	var slice []int = []int {1,1}
	if n == 1 || n == 2{
		return slice[:n] 
	}

	for i:=2; i<n; i++ {
		slice = append(slice, slice[i-1]+slice[i-2])
	}
	return slice
}

func divided_sorted(arr [7]int, left_index int, right_index int, value int) {
	middle := (left_index + right_index)/2
	if left_index > right_index {
		fmt.Println("not exist")
		return
	}
	if arr[middle] > value {
		right_index = middle - 1
	    divided_sorted(arr, left_index, right_index, value)
	} else if arr[middle] < value{
		left_index = middle + 1
		divided_sorted(arr, left_index, right_index, value)
	} else {
		fmt.Println("index= ", middle)
	}
}


func modifyUser(user_info map[string]map[string]string, user_name string) {
	user, ok := user_info[user_name]
	if ok{
		user["pwd"] = "88888888"
	} else {
		// new_user := map[string]string{
		// 	"pwd": "8888888",
		// 	"nickname": user_name,
		// }
		// user_info = append(user_info, new_user)
		user_info[user_name] = make(map[string]string, 2)
		user_info[user_name]["pwd"] = "88888888"
		user_info[user_name]["nickname"] = user_name
		
	}
}


type Circle struct {
	Radius float32
}

func (c Circle) area() float32 {
	return 3.14 * c.Radius * c.Radius
}
func main () {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintln(writer, "go server start succed")
	})
	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}