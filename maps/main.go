package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["red"] = "f0000"
	printMap(m)
	fmt.Println(m)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("The key is", key, "and the value is", value)
	}
}
