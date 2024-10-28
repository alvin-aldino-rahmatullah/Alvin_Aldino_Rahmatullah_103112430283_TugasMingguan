<--SOAL 2 MINGGU 3-->

Mahasiswa dikatakan cumlaude apabila lulus maksimal 8 semester dan nilai EPrT tidak kurang dari 500
Masukan : terdiri dari dua nilai. Nilai pertama adalahnya jumlah semester yang ditempuh hingga lulus, dan nilai kedua adalah skor EPrT.
Keluaran : adalah boolean yang menyatakan true apabila mahasiswa lulus dengan predikat cumlaude, atau false apabila sebaliknya.
_________________________________________________________________________________________________________________________________________
<--NOTASI PSEUDOCODE-->

Program : cumlaude
{Program ini dibuat untuk mempermudah seleksi seseorang apakah lulus cumlaude dengan hal yang sudah di tentukan seperti lulus di bawah 8 semester dan mendapatkan skor EPrT lebih dari 500 jika iya maka program akan menyatakan mahasiswa tersebut lulus atau true}

Rumus : 
{Memakai tipe data int karena menghitung bilangan}
{Membuat Variabel yang dibutuhkan}
var semester, eprtScore int
{Memakai && untuk menyatakan bahwa kedua perbandingan harus sama sama true jika tidak maka akan false}

Algoritma :
{User harus memasukkan jumlah score dan semester}
  input (semester)
  input (eprtScore)
{Mentukan apakah semester lebih kecil dari 8 dan score lebih besar dari 500 maka && yaitu harus true dari semester dan eprtScore}
  cumlaude :=  Cumlaude := (semester <= 8) && (eprtScore >= 500)
{Hasil cek dari cumlaude}
  output : (cumlaude)
end program
_______________________________________________________________________________________________________________________________________
<--NOTASI BAHASA PEMOGRAMAN GOLANG-->

package main

import "fmt"

func main() {
    /*Membuat variabel semester dan eprtScore lalu int sebagai tipe data karena memakai bilangan bulat*/
    var semester, eprtScore int

    /*Program meminta user untuk memasukkan semester yang sudah di tempuh*/
    fmt.Print("Masukkan jumlah semester yang ditempuh: ")
    /*Program Scan berguna untuk menyimpan bilangan yang di masukkan user kedalam variabel semester*/
    fmt.Scan(&semester)
    /*Meminta user untuk memasukkan  yang sudah di tempuh*/
    fmt.Print("Masukkan skor EPrT: ")
    /*Program Scan berguna untuk menyimpan bilangan yang di masukkan user kedalam variabel eprtScore*/
    fmt.Scan(&eprtScore)

    /*Memeriksa apakah semesternya lebih kecil dari pada 8 dan memeriksa apakah score eprt lebih besar dari 500*/
    Cumlaude := (semester <= 8) && (eprtScore >= 500)

   /*Output dari cek yg di lakukan cumlaude*/
    fmt.Println("Predikat Cumlaude:", Cumlaude)
}


________________________________________________________________________________________________________________________________________
<--CONTOH SAAT MENJALANKAN PROGRAM-->

contoh true :
input :
Masukkan jumlah semester yang ditempuh: 7
Masukkan skor EPrT: 520
output :
Predikat Cumlaude: true

Contoh False :
input :
Masukkan jumlah semester yang ditempuh: 10
Masukkan skor EPrT: 574
output :
Predikat Cumlaude: false
"false terjadi karena semester melebihi 10 walaupun mendapatkan score yang tinggi"