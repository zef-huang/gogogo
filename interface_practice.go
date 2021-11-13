// 练习 golang 的方法和接口特性
package main


import (
	"fmt";
	"strconv";
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




func main()  {
	ips := map[string]IPAddr {
		"test1": {1, 2, 3, 4},
		"test2": {8, 8, 8, 8},
	}

	for k, v := range(ips){
		fmt.Println(k, v)
	}
}