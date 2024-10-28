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