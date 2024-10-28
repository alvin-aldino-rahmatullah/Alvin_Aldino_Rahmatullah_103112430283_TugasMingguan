<--SOAL 4 MINGGU 2-->

Sebuah program digunakan untuk menentukan tiga digit nilai yang terdapat pada suatu bilangan bulat positif x.
Masukan berupa bilangan bulat positif x yang kurang atau sama dengan 999.
Keluaran terdiri dari dari tiga bilangan d1, d2, dan d3 yang menyatakan digit pertama, kedua dan ketiga dari x.
Petunjuk: satuan dapat diperoleh apabila bilangan apapun dimodulo dengan 10
__________________________________________________________________________________________________________________
<--NOTASI PSEUDOCODE-->

Program : Digit
{Program ini dibuat untuk menghitung setiap digit pada bilangan positif dari 1 sampai 999, maka dari itu kita memakai konsep modulo agar bisa melihat sisa bagi setiap bilangan untuk mendapatkan output yang diingingkan }

rumus :
{memakai uint untuk menghitung bilangan bulat positif}
{x adalah input dari user}
{Cara menghitung setiap digit}
Untuk Mencari digit pertama (x / 100)
untuk mencari digit ke dua (x % 100)/10
untuk mencari digit ke tiga (x % 10 )

Algoritma:
{input berguna sebagai tempat user memasukkan nilai x}
input (x)
{Cara menghitung digit 1}
  d1 = x / 100
{Cara menghitung digit 2}
 d2 = (x % 100) / 10
{Cara menghitung digit 3}
  d3 = x % 10
{output akan keluar dari sini dengan variabel yang sudah di hitung tadi pada d1, d2, dan d3}
  output = (digit1 = 1 ,digit2 = d2, digit 3= d3)
endprgram
__________________________________________________________________________________________________
<--NOTASI BAHASA PEMOGRAMAN GOLANG-->

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

________________________________________________________________________________________________________________________________________
<--CONTOH SAAT MENJALANKAN PROGRAM-->

input 
Masukkan bilangan bulat positif : 123
Output
Digit pertama: 1
Digit kedua: 2
Digit ketiga: 3