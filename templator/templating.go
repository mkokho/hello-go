package main

import (
	"text/template"
	"fmt"
	"os"
)

func main() {

	markerNames := []string{
		"xx",
		"x>x",
		"x&x",
		"x_x",
		"x'x",
		"tưởng",
		"tưởng",
	}

	for _, m := range markerNames {
		body := fmt.Sprintf("␞{.%s}␞", m)
		values := map[string]string {m: "value"}
		tmpl, err := template.New("test").Parse(body)

  	fmt.Printf("Body is: %s", body)
		if err != nil {
			fmt.Print(" failed to parse: ")
			fmt.Print(err.Error())
		} else {
			fmt.Print(" Result is: ")
			err = tmpl.Execute(os.Stdout, values)
		}
		fmt.Print("\n")
	}

}
