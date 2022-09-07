package main

import (
    "fmt"
)

type Animal struct {
    name string
    age int64
}


func (a *Animal) DrinkWater() {
    fmt.Println(a.name, "drink water")
}


func main() {
    // 封装
    cat := Animal{"cat", 2}
    cat.DrinkWater()

    // 继承

    // 多态
}