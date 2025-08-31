package main

import (
	"fmt"
	"sync"
)

func tryMutex() {
	var wg sync.WaitGroup
	var mu sync.Mutex   // "Spidol" kita
	var counter int = 0 // Variabel yang dilindungi

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			
			mu.Lock()   // 1. Ambil spidol sebelum menulis
			counter++
			mu.Unlock() // 2. Kembalikan spidol setelah selesai
		}()
	}

	wg.Wait()
	// Hasilnya akan selalu konsisten 1000.
	fmt.Println("Hasil akhir (benar):", counter)
}