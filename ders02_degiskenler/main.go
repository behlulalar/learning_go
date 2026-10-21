package main

import "fmt"

func main() {

	var isim string = "Behlul"
	var yas int = 22
	var boy float64 = 1.72
	var ogrenciMi bool = true

	fmt.Printf("Ad: %s, \n Yaş: %d, \n Boy:%.2f, \n Öğrenci Mi: %t\n", isim, yas, boy, ogrenciMi)

	isimm := "Behlul"
	yass := 22

	fmt.Printf("Ad: %s, \n Yaş: %d, \n", isimm, yass)

	const name = "Behlul"
	fmt.Printf("Name: %s, \n", name)

	var sayac int
	fmt.Println(sayac)

	toplam := yas + boy
	fmt.Println(toplam)

}
