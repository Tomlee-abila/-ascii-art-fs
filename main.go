package main

import (
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) != 3{
		fmt.Println("Only two arguments are accepted after main.go.\nUsage: go run . [STRING] [BANNER]\nEX: go run . something standard")
	}
	
}