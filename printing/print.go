package main

import (
	"fmt"
	"os"
)

func main() {
	token := os.Getenv("TERRATEST_GITHUB_TOKEN")
	fmt.Print(token)
	fmt.Print("test")

}
