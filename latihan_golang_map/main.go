package main

import "fmt"

func main() {
	data := map[string]interface{}{
		"data1": 12,
		"data2": "data2",
	}

	fmt.Println(data["data1"])
}
