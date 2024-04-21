package main

import "fmt"

// Run - is going to be responsible for
// instantiation and startup of our
// go application
func Run() error {
	fmt.Println("starting up our application")
	return nil
}
func main() {
	fmt.Println("Go REST API Course")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
