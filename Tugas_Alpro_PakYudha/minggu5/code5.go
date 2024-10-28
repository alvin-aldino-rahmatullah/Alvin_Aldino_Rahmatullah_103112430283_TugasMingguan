package main

import "fmt"

func main () {
	/*Membuat variabel n dan tipe data uint untuk menentukan bilangan bulat positif*/
    var n uint
	/*Program meminta user untuk memasukkan nilai bilangan positif*/
    fmt.Print("Masukkan bilangan positif: ")
	/*Program Scan berguna untuk menyimpan bilangan positif yang di masukkan user kedalam variabel n*/
    fmt.Scan(&n)

	/*looping atau for bilangan 1 sampai n untuk mencari faktor*/
    for i := uint(1); i <= n; i++ {
		/*Output yang di keluarkan adalah input n yang di modulo dengan hasil looping dan hasil dari modulo harus 0, jika hasil sisa 0 maka true*/
        fmt.Println(i, n % i == 0)
    }
}
