package main

import (
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) > 3 || len(os.Args) < 2{
		fmt.Println("Only two arguments are accepted after main.go.\nUsage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}

	BannerFile := ".txt"
	if len(os.Args) > 2{
		BannerFile = os.Args[2]+BannerFile
	}else{
		BannerFile = "standard"+BannerFile
	}
}