package main

import "fmt"

func main () {
	/*Membuat Variabel x, d1, d2, d3,  dan menggunakan tipe data int untuk menghitung bilangan bulat*/
    var x, d1, d2, d3 uint

	/*"Program meminta user untuk memasukkan bilangan bulat positif*/
    fmt.Print("Masukkan bilangan bulat positif : ")
    /*Program Scan berguna untuk menyimpan bilangan yang di masukkan user kedalam variabel x*/
    fmt.Scan(&x)

	/*Cara menghitung digit 1 lalu d1 akan menyimpan data hasil tersebut*/
    d1 = x / 100        
	/*Cara menghitung digit 2 lalu d2 akan menyimpan data hasil tersebut*/
    d2 = (x % 100) / 10 
	/*Cara menghitung digit 3 lalu d3 akan menyimpan data hasil tersebut*/
    d3 = x % 10      

	/*Lalu output akan keluar dari sini dengan variabel yang sudah di hitung tadi pada d1, d2, dan d3*/
    fmt.Println("Digit pertama:", d1)
    fmt.Println("Digit kedua:", d2)
    fmt.Println("Digit ketiga:", d3)
}