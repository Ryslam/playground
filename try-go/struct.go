package main

import "fmt"

// Mendefinisikan tipe data baru bernama 'Mahasiswa'
type Mahasiswa struct {
    Nama    string
    NIM     string
    Angkatan int
    Aktif   bool
}

func tryStruct() {
    // Membuat variabel 'mhs1' dengan tipe data 'Mahasiswa'
    mhs1 := Mahasiswa{
        Nama:    "Budi Santoso",
        NIM:     "13521001",
        Angkatan: 2021,
        Aktif:   true,
    }

    fmt.Println(mhs1)
	
}