package main

import(
	"fmt"
)

type Bentuk interface{
	keliling() float64
	luas() float64
}

type persegipjg struct{
	panjang, lebar float64
}

func (pj persegipjg) keliling() float64 {
	return 2*pj.panjang + 2*pj.lebar
}

func (pj persegipjg) luas() float64 {
	return pj.panjang * pj.lebar
}


func main()  {
	kotak1 := persegipjg{5,10}
	hitungbangunan(kotak1)
}

func hitungbangunan(b Bentuk)  {
	fmt.Println("Keliling", b.keliling())
	fmt.Println("Luas :", b.luas())
}