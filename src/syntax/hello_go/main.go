package main

import "fmt"

func main() {
	DeferClosureLoopV1()
}

func DeferClosureLoopV1() {
	var j int
	for i := 0; i < 10; i++ {
		j = i
		defer func() {
			fmt.Println(j)
		}()
	}
}
