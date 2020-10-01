package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	data, e := ioutil.ReadFile("names.txt")

	if e != nil {
		fmt.Println(e)
		return
	}

	nameSlice := make([]Name, 0)
	strData := string(data)
	names := strings.Split(strData, "\n")
	for _, name := range names {
		fullName := strings.Split(name, " ")
		fName := fullName[0]
		lName := fullName[1]

		nameSlice = append(nameSlice, Name{fName, lName})
	}

	for _, name := range nameSlice {
		fmt.Println("First name: ", name.fname, "Last name: ", name.lname)
	}
}
