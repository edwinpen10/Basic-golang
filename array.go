package main

import (
	"fmt"
)

func main()  {
	var member [3]string

	member[0] = "edwin"
	member[1] = "tokichi"
	member[2] = "retstu"

	angka := [...]int{1,2,3,50,90,900,893}
	angka[1]=200

	// angka := [3]int{1,2,3}
	// angka[1]=200

	fmt.Println(angka)
	fmt.Printf("------------Lopping data member------------\n")
	for i := 0; i < len(member); i++ {
		fmt.Println(member[i])
	}

	hari := [...]string{"senin","selasa","rabu","kamis","jumat","sabtu","minggu"}
	fmt.Printf("--------------Looping data hari-----------\n")
	for index, val := range hari {
		
		fmt.Printf("Hari %d = %s\n", index, val)
	}
}