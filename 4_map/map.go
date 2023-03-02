package main

import "fmt"

func main() {

	data := map[string]string{
		"data1": "value_data_1",
		"data2": "value_data_2",
		"data3": "value_data_3",
	}

	fmt.Println(data["data1"])

	// for untuk access data map

	for key, value := range data {
		fmt.Printf("Key : %s , data : %s \n", key, value)
	}
}
