package main

import ("fmt")

func main()  {
	member := [5]string {"naruto","sasuke","sakura","hinata","boruto"}
	
	sliceMembebr := member[:]
	sliceMembebr[1] = "yamato"

	fmt.Println(member)
	fmt.Println(sliceMembebr)

	fmt.Println(len(sliceMembebr)) //menampilkan panjang array
	fmt.Println(cap(sliceMembebr)) //menampilkan kapasistas didalam array

	human := make([]string, 5, 10)
	
	fmt.Println(human)
	fmt.Println(len(human))
	fmt.Println(cap(human))

	
	humanity :=  make([]string, 5)
	copy(humanity, sliceMembebr)

	humanity = append(humanity,"jiraya","orocimaru")

	fmt.Println(humanity)
}