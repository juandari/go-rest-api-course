package main

import "fmt"

func Run() error {
	fmt.Println("starting our app")

	return nil
}

func main() {
	fmt.Println("Go REST API Course")

	if err:= Run(); err != nil {
		fmt.Println(err)
	}
}