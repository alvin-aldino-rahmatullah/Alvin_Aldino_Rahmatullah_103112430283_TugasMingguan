<--SOAL 9 MINGGU KE 5-->

Buatlah program dalam bahasa Go untuk menampilkan faktor dari suatu bilangan. Faktor adalah bilangan yang habis membagi suatu bilangan. Contoh: Faktor dari 15 adalah 1, 3, 5 dan 15. Faktor dari 24 adalah 1, 2, 3, 4, 6, 8, 12, dan 24.
Masukan terdiri dari sebuah bilangan bulat positif N.
______________________________________________________________________________________________________________________________________________________________________
<--NOTASI PSEUDOCODE-->

Program : Faktor Bilangan
{Program ini bisa menentukan setiap faktor atau bilangan yang habis di bagi dengan bilang dengan menggunakan bahasa pemograman golang. Dalam program ini bisa menggunaka modulo agar bisa mengetahui apakah setiap loop tersebut bisa habis jika di modulo}

Rumus :
{Memakai Tipe data uint untuk melihat bilangan bulat yang akan dimasukkan bersifat positif}
{n adalah input dari user atau bilangan bulat positif yang di masukkan user}
{Code untuk mengetahui modulo}
 (n % i == 0)

algoritma :
{input berguna sebagai tempat user memasukkan nilai n}
 input (n)
{Proses looping bilangan 1 sampai n}
for i := uint(1); i <= n; i++ 
       {Output yang di keluarkan adalah input n yang di modulo dengan hasil looping dan hasil dari modulo harus 0, jika hasil sisa 0 maka true}
        output(i, n % i == 0)
    endfor
endprogram
_______________________________________________________________________________________________________________________________________________________________________
<--NOTASI BAHASA PEMOGRAMAN GOLANG-->

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

________________________________________________________________________________________________________________________________________________________________________
<--CONTOH SAAT MENJALANKAN PROGRAM-->

input :
Masukkan bilangan positif: 5
output :
1 true 
2 false
3 false
4 false
5 true 