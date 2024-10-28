<--SOAL 4 MINGGU 4-->

Sebuah program digunakan untuk menyimpan data karyawan yang berisi nama, gaji pokok, tunjangan, potongan, dan total gaji.
Masukan: terdiri dari satu string dan tiga bilangan riil yang menyatakan nama, gaji pokok, tunjangan, dan potongan.
Keluaran: terdiri dari satu bilangan riil yang menyatakan total gaji karyawan.
Catatan: Gunakan tipe bentukan untuk menyimpan data karyawan tersebut.
______________________________________________________________________________________________________________________________
<--NOTASI PSEUDOCODE-->

Program : Gaji Karyawan
{Program ini dapat mengeluarkan output berupa nama, gaji pokok, tungangan, potongan dan menjumlahkan semuanya menjadi total gaji, sehingga user hanya perlu memasukkan nama dan beberapa input yang di butuhkan lalu program otomatis menghitung total gaji karyawan dengan cara gajiPokok + tunjangan - potongan}

Rumus :
{Membuat variabel nama dengan tipe data string untuk menyimpan nama karyawan}
{Membuat variabel gajiPokok, tunjangan, potongan  dengan tipe data float64 }
{menghitung totalGaji}
gajiPokok + tunjangan - potongan

algoritma :
{input berguna untuk user memasukkan nama, gaji pokok, tunjangan dan potongan sebelum semuanya di jumlahkan}
input (nama)
input (gajiPokok)
input (tunjangan)
input (potongan)
{Proses menentukan total gaji yang akan di berikan ke karyawan}
totalGaji := gajiPokok + tunjangan - potongan
{Output berupa hasil dari penjumlahan dan pengurangan dari total gaji}
output = (nama, gajiPokok, tunjangan, potongan, totalGaji) 


________________________________________________________________________________________________________________________________
<--NOTASI BAHASA PEMOGRAMAN GOLANG-->

package main

import "fmt"


func main () {
    /*Membuat variabel nama dengan tipe data string*/
    var nama string
    /*Membuat variabel nama dengan tipe data float64*/
	var gajiPokok, tunjangan, potongan  float64

    /*Program meminta user untuk memasukkan nama karyawan*/
    fmt.Print("Masukkan nama karyawan: ")
    /*Program Scan berguna untuk menyimpan nama yang di masukkan user kedalam variabel*/
    fmt.Scanln(&nama)

    /*Program meminta user untuk memasukkan gaji pokok karyawan*/
    fmt.Print("Masukkan gaji pokok: Rp ")
    /*Program Scan berguna untuk menyimpan gaji Pokok yang di masukkan user kedalam variabel*/
    fmt.Scanln(&gajiPokok)

    /*Program meminta user untuk memasukkan tunjanga karyawan*/
    fmt.Print("Masukkan tunjangan: Rp ")
    /*Program Scan berguna untuk menyimpan tunjangan yang di masukkan user kedalam variabel*/
    fmt.Scanln(&tunjangan)

    /*Program meminta user untuk memasukkan potongan karyawan*/
    fmt.Print("Masukkan potongan: Rp ")
    /*Program Scan berguna untuk menyimpan potongan yang di masukkan user kedalam variabel*/
    fmt.Scanln(&potongan)

    /*Proses menghitung total gaji*/
    totalGaji := gajiPokok + tunjangan - potongan


    /*Output yang dikeluarkan setelah mendapatkan total gaji dan beberapa informasi yang telah tersimpan dalam variabel*/
    fmt.Println("INFORMASI KARYAWAN")
    fmt.Println("Nama:", nama)
    fmt.Printf("Gaji Pokok: Rp %.2f\n", gajiPokok)
    fmt.Printf("Tunjangan: Rp %.2f\n", tunjangan)
    fmt.Printf("Potongan: Rp %.2f\n", potongan)
    fmt.Printf("Total Gaji: Rp %.2f\n", totalGaji)
}

_____________________________________________________________________________________________________________________________________________
<--CONTOH SAAT MENJALANKAN PROGRAM-->
Input: 
Masukkan nama karyawan: budi
Masukkan gaji pokok: Rp 5000000.00
Masukkan tunjangan: Rp 1000000.00
Masukkan potongan: Rp 500000.00

output:
INFORMASI KARYAWAN
Nama: budi
Gaji Pokok: Rp 5000000.00
Tunjangan: Rp 1000000.00 
Potongan: Rp 500000.00   
Total Gaji: Rp 5500000.00
