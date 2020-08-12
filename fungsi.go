package main

import (
	"fmt"
	"strconv"
)


func tambah(a int, b int) int{
	 hasil := a+b
	 return hasil
}

func biodata(umur int, nama string, sekolah string)  string{
	umurku :=  strconv.Itoa(umur)
	return nama +", umur saya adalah  "+ umurku + "dan sekolah saya di "+sekolah
}

func siswa(kelas int,  nama string, sekolah string)  (string, int){
	
	return "nama saya "+ nama +" sekolah saya di "+sekolah,
	kelas
}

func siswa2(kelas int,  nama string, sekolah string)  (string, string){
	kelasku :=  strconv.Itoa(kelas)
	return "nama saya "+ nama +" sekolah saya di "+sekolah,
	"dan saya kelas "+kelasku
}

func guru(nama string, mapel string)  (namaguru string, mapelku string){
		namaguru= nama
		mapelku= mapel
		return 
}


func main()  {
	//fmt.Println(tambah(2,10))
	//fmt.Println(biodata(24,"edwin","SMKN1 Pangkalanbun"))

	//infoBasic, classInfo := siswa(3, "edwin", "SMKN1 Pangkalanbun")
	infoBasic2, classInfo2 := siswa2(3, "edwin", "SMKN1 Pangkalanbun")
	fmt.Println(infoBasic2, classInfo2)

	infoBio, mapelInfo := guru("steve job","swift code for apple product")

	fmt.Println(infoBio, mapelInfo)
}