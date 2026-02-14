package main

import "fmt"

func Hello(Name string) string {
	return "Hello, " + Name
}

func main() {
	fmt.Println(Hello("Chris"))
}
