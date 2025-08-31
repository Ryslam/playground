package main

import (
	"fmt"
	"sync"
	"time"
)

func pekerja(id int, wg *sync.WaitGroup) {
	// defer akan memastikan wg.Done() dipanggil di akhir
	defer wg.Done()

	fmt.Printf("Pekerja %d memulai tugas\n", id)
	time.Sleep(1 * time.Second) // Simulasi pekerjaan
	fmt.Printf("Pekerja %d selesai\n", id)
}

func tryWaitGroup() {
	var wg sync.WaitGroup

	// Kita akan menjalankan 3 Goroutine "pekerja"
	for i := 1; i <= 3; i++ {
		// Manajer: "Saya punya 1 tugas baru"
		wg.Add(1)
		go pekerja(i, &wg)
	}

	fmt.Println("Fungsi main menunggu semua pekerja selesai...")
	// Manajer: "Saya tunggu sampai counter jadi nol"
	wg.Wait()

	fmt.Println("Semua pekerjaan selesai. Program berakhir.")
}