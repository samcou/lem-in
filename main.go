package main

import (
	"fmt"
	"lemin/lemin"
)

func main() {
	data, err := lemin.ReadData()
	if err != nil {
		fmt.Println(err)
		return
	}
	lemin.PrintOutput(data)
}
