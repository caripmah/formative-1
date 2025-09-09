package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Soal 1
func soal1() {
	kata1 := "Bootcamp"
	kata2 := "Digital"
	kata3 := "Skill"
	kata4 := "Sanbercode"
	kata5 := "Golang"
	fmt.Println("Soal 1:", kata1, kata2, kata3, kata4, kata5)
}

// Soal 2
func soal2() {
	halo := "Halo Dunia"
	hasil := strings.Replace(halo, "Dunia", "Golang", 1)
	fmt.Println("Soal 2:", hasil)
}

// Soal 3
func soal3() {
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kalimat := kataPertama + " " +
		strings.Title(kataKedua) + " " +
		kataKetiga[:len(kataKetiga)-1] + strings.ToUpper(kataKetiga[len(kataKetiga)-1:]) + " " +
		strings.ToUpper(kataKeempat)

	fmt.Println("Soal 3:", kalimat)
}

// Soal 4
func soal4() {
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	a1, _ := strconv.Atoi(angkaPertama)
	a2, _ := strconv.Atoi(angkaKedua)
	a3, _ := strconv.Atoi(angkaKetiga)
	a4, _ := strconv.Atoi(angkaKeempat)

	jumlah := a1 + a2 + a3 + a4
	fmt.Println("Soal 4: Hasil penjumlahan =", jumlah)
}

// Soal 5
func soal5() {
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	
	panjang, _ := strconv.Atoi(panjangPersegiPanjang)
	lebar, _ := strconv.Atoi(lebarPersegiPanjang)
	alas, _ := strconv.Atoi(alasSegitiga)
	tinggi, _ := strconv.Atoi(tinggiSegitiga)


	var luasPersegiPanjang int = panjang * lebar
	var kelilingPersegiPanjang int = 2 * (panjang + lebar)
	var luasSegitiga int = (alas * tinggi) / 2


	fmt.Println("Soal 5:")
	fmt.Println("Luas Persegi Panjang =", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang =", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga =", luasSegitiga)
}

func main() {
	soal1()
	soal2()
	soal3()
	soal4()
	soal5()
}