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

