package main


import(
	"fmt"
)

func main(){
	fmt.Println("fungsi ini menghasilkan ", jumlahAngka(9, 10))
	fmt.Println("fungsi ini menghasilkan ", kurangAngka(19, 10))
	fmt.Println("fungsi ini menghasilkan ", bagiAngka(20, 10))
}

func printText() int{
	fmt.Println("Mencoba print hasil ini ")
	return 100
}

func jumlahAngka(angka1 int, angka2 int) int{
	return angka1 + angka2
}

func kurangAngka(angka1 int, angka2 int) int{
	return angka1 - angka2
}

func bagiAngka(angka1 int, angka2 int)int{
	return angka1 / angka2
}