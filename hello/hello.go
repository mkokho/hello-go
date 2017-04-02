package main

import "fmt"

func main() {
	raw := make(map[string]interface{})
	raw["x"] = 1

	html, ok := raw["x"].(string)
	if !ok {
			fmt.Printf("unexpected type for html template")
	}
	fmt.Printf("%s", html)
}
