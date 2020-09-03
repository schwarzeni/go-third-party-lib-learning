package main

import (
	"errors"
	"fmt"
	"go-rafaeldias-async-play/myasync"
	"time"
)

func fib(p, c int) (int, int) {
	return c, p + c
}

func main() {

	//// execution in series.
	//res, e := async.Waterfall(async.Tasks{
	//	fib,
	//	fib,
	//	fib,
	//	func(p, c int) (int, error) {
	//		return c, nil
	//	},
	//}, 0, 1)
	//
	//if e != nil {
	//	fmt.Printf("Error executing a Waterfall (%s)\n", e.Error())
	//	return
	//}
	//
	//fmt.Println(res[0].(int)) // Prints 3

	// execution in series.
	//res, _ := myasync.Waterfall(myasync.Tasks{
	//	func(arr []int) (res []int) {
	//		for i := range arr {
	//			if arr[i] > 3 {
	//				res = append(res, arr[i])
	//			}
	//		}
	//		return
	//	},
	//	func(arr []int) (res []int) {
	//		for i := range arr {
	//			res = append(res, arr[i]+2)
	//		}
	//		return
	//	},
	//}, []int{1, 2, 3, 4, 5, 6})
	//
	//fmt.Println(res[0].([]int))

	res := myasync.Concurrent(myasync.MapTasks{
		"one": func() int {
			for i := 'a'; i < 'a'+26; i++ {
				fmt.Printf("%c ", i)
			}

			return 1
		},
		"two": func() int {
			time.Sleep(time.Second)
			for i := 0; i < 27; i++ {
				fmt.Printf("%d ", i)
			}

			return 2
		},
		"three": func() error {
			for i := 'z'; i >= 'a'; i-- {
				fmt.Printf("%c ", i)
			}

			return errors.New("self define error in three")
		},
	})

	fmt.Printf("Results from task 'two': %v\n", res.Value("two"))
	fmt.Printf("Error from task 'three': %+v", res.Error("three"))

}
