package main

import (
	"fmt"
	"goasync"
	"time"
)

func task() goasync.Async {
	return goasync.Task(func() (interface{}, error) {
		fmt.Println("async begin")
		time.Sleep(1 * time.Second)
		fmt.Println("async end")
		return "ok", nil
	})
}

func main() {
	fmt.Println("main begin")

	task := task()
	fmt.Println("await begin")
	result, _ := task.Await()
	fmt.Println("await end")
	fmt.Println(result)

	fmt.Println("main end")
}
