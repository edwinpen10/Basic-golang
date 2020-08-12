package main

import(
	"fmt"
	"strconv"
)

func main()  {
	//int to string
	umur := 12
	fmt.Println("umur ku adalah "+ strconv.Itoa(umur))

	//string to int
	gaji := "20000"
	gaji2,_:=strconv.Atoi(gaji)
	bonus:=gaji2+4500

	fmt.Println(bonus)
}