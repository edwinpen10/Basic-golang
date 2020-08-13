package main

import ("fmt")

func main()  {
	member := [5]string {"naruto","sasuke","sakura","hinata","boruto"}
	
	sliceMembebr := member[:]
	sliceMembebr[1] = "yamato"

	fmt.Println(member)
	fmt.Println(sliceMembebr)
}