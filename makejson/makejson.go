package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string

	fmt.Printf("Enter your name: ")
	_, _ = fmt.Scan(&name)

	fmt.Printf("Enter your address: ")
	_, _ = fmt.Scan(&address)

	pMap := map[string]string{"name": name, "address": address}

	pJSON, e := json.Marshal(pMap)

	if e == nil {
		fmt.Println("JSON: ", pJSON)

		var pMap2 map[string]string
		e2 := json.Unmarshal(pJSON, &pMap2)

		if e2 == nil {
			fmt.Println("Map: ", pMap2)
		} else {
			fmt.Println(e2)
		}
	} else {
		fmt.Println(e)
	}
}
