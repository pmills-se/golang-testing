package main

import (
	"fmt"
	"strconv"
)

func main() {
	thresh := "20.00"
	_, threshold := strconv.ParseFloat(thresh, 10)
	fmt.Println(threshold)

}
