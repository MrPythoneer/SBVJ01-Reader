package main

import (
	"fmt"
	"joshua/green/starbound/sbvj"
)

func main() {
	data, _ := sbvj.ReadFile("file.player")

	fmt.Println("Name:", data.Name)
	fmt.Printf("Versioned (%t) = %d\n", data.Versioned, data.Version)
	fmt.Println(data.Value)
}
