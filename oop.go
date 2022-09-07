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


type Bird struct {
    Animal
}


func main() {
    // 封装
    cat := Animal{"cat", 2}
    cat.DrinkWater()

    // 继承: 通过组合的方式实现
    bird := Bird{Animal{"bird", 1}}
    bird.DrinkWater()

    // 多态
}