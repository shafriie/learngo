package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	// nameInfo, kelasInfo := getData("Yusuf", "IK 2 1")

	// fmt.Println(nameInfo, kelasInfo)

	// detailInfo, umurInfo := getJurusan("Informatika", "Syafrie", 23)

	// fmt.Println(detailInfo)

	// fmt.Println(umurInfo)

	// a, b := getLokasi("Fatahillah", "Jakarta Barat", 13)

	// fmt.Println(a)
	// fmt.Println("Nomor rumah nya adalah ", b)

	fmt.Println(printText())
}

func printText() string{
	return "Hello world, i like golang, and i want to become Golang Developer"
}

func getLokasi(namaTempat string, daerah string, nomorRumah int) (detailAlamat string, no int){
	detailAlamat = "Tempatnya adalah " + namaTempat + " yg terletak di " + daerah	
	no = nomorRumah
	return
}

func getBio(age int, name string, status string) string{
	ageNow := strconv.Itoa(age)
	return name + " adalah seorang " + status + " saat ini berumur " + ageNow + " tahun"
}

func getData(name string, kelas string) (string, string){
	return strings.TrimSpace("Nama saya adalah " + name),
			",dan dari kelas " + kelas
}

func getJurusan(jurusan string, nama string, umur int) (string, int){
	return1 := "Nama nya adalah " + nama + " dari jurusan " + jurusan + " umurnya sekarang "
	return2 := umur
	return return1, return2
}

