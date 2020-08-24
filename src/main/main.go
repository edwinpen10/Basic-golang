package main

import (
	"fmt"
	"go_learning/tutorial_1/src/models"
)

func main()  {
	user := models.Member{
		Name: "edwin",
		Age: 20
	}

	fmt.Println(user)
}