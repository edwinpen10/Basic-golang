package main

import(
	"fmt"
)

type Member struct{
	name string
	age int
}

// func (m Member) getInfo() {
// 	fmt.Println("hallo edwin")
// }

//Dengan pointer
// func (m *Member) getInfo() {
// 	m.name = "penangsang"
// }

//Dengan fungsi
func (m *Member) changeName(newName string, newAge int) {
	m.name = newName
	m.age = newAge
}


func main()  {
	user:=Member{"edwin",24}
	user.changeName("markzukerberg", 50)
	// user.getInfo()
	fmt.Println(user)
}