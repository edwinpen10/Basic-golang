package main

import "fmt"

func main()  {
	var point = 10

	switch  {

	case point > 10 : fmt.Println("lulus dengan nilai sempurna")

	case point == 8 : fmt.Println("lulus dengan nilai bagus")

	case point <=5 : fmt.Println("Mengulang")

	default : fmt.Println("Mengulang lagi")
		
	}

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}


}
