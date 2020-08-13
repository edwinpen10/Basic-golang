package main

import (
	"fmt"
)

func main()  {
	hewan := [3][2]string {
		{"golang", "php"},
		{"Javascript", "C#"},
		{"Java", "Ruby"},
	}
	fmt.Println(hewan[0][1])

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(hewan[i][j])
		}
	}

	for h := 0; h < len(hewan); h++ {
		for k := 0; k < len(hewan[h]); k++ {
			fmt.Println(hewan[h][k])
		}
	}
}