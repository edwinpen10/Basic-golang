package main

import (
	"fmt"
)

var member = map[int]string{
	332421233 : "Edwin",
	123123331 : "mantull",
}

func main()  {
	
	

	var newMember = member

	newMember[332421233]="okey coy"

	fmt.Println(member[332421233])
	fmt.Println(member)

	

	for id, name := range member {
		fmt.Println(id, name)
	}

	delete(member, 332421233) //fungsi untuk menghapus data

	if checkMmeber(332421233) {

		fmt.Println("data is already")
	}else{
		fmt.Println("data not found")
	}
}

func checkMmeber(id int) bool {
	_, exitis := member[id]
	return exitis;
}
