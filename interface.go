package main

import(
	"fmt"
	"math"
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

type lingkaran struct{
	radius float64
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.radius * l.radius
}

func (l lingkaran) luas() float64  {
	
	return 2 * math.Pi * l.radius
}

func main()  {
	kotak1 := persegipjg{5,10}
	hitungbangunan(kotak1)
	lingkaran1 := lingkaran{10}
	hitungbangunan(lingkaran1) 
}

func hitungbangunan(b Bentuk)  {
	fmt.Println("Keliling", b.keliling())
	fmt.Println("Luas :", b.luas())
}