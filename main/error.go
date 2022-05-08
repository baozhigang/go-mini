package main

import (
	"errors"
	"fmt"
	"log"
)

// log.Fatalf 适用于比较严重的错误，影响了程序的执行。比如连接数据库，出现问题直接退出程序。
// fmt.Println(err) return 适用于用户交互层面，出现问题输出信息给用户，退出或者不退出看情况
// errors.New() 类似PHP的抛出异常
func main() {
	age, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
		return // 上线后若有报错，直接退出程序
	}

	name, id := "zhangsan", age
	err2 := fmt.Errorf("user %s id %d not found", name, id)
	if err2 != nil {
		fmt.Println(err2)
		return

		log.Fatalf("user %s id %d not found", name, id) // 输出信息，并退出程序，
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("wrong number") // 子方法抛异常
	}

	return f / 2, nil
}
