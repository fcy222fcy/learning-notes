package main

import "fmt"

func main() {
	//panic相关

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	defer func() {
		fmt.Println("This will be printed")
	}()
	panic("Something went wrong")

	defer func() {
		fmt.Println("This won't be printed")
	}()
	
}
