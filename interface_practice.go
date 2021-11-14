// 练习 golang 的方法和接口特性
package main


import (
	"fmt";
	"strconv";
	"math";
	// "reflect"
)


// 方法练习，将 4 个值的数组转为 ip 地址
type IPAddr [4]int

func (ip IPAddr) String() string {
	ip_string := ""
	for i, v := range(ip) {
		v_string := strconv.Itoa(v)
		ip_string = ip_string + v_string
		if (i == 3){
			break
		}
		ip_string = ip_string + string('.')
	} 

	return ip_string
}


// 报错方法练习
// error 是内部接口需要对应的 Error() 方法
type SqrtNegativeErr float64


func (e SqrtNegativeErr) Error() (string){
	return fmt.Sprint(float64(e), " cannot have sqrt root")
}


// 平方根函数
func get_square_root(x float64) (float64, error){
	left := 0.0
	right := x

	if (x < 0) {
		return 0.0, SqrtNegativeErr(x)
	}

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

	return right, nil
}



func main()  {
	_, value := get_square_root(-1)
	// value := SqrtNegativeErr(-1.0)
	fmt.Println(value)
	// fmt.Println(ok)
}