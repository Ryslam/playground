package main

import (
	"fmt"
	"time"
)

func katakan(pesan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(pesan)
		time.Sleep(500 * time.Millisecond)
	}
}

func tryGoroutines() {
	// Menjalankan fungsi ini secara normal (blocking)
	go katakan("Dunia")

	// Menjalankan fungsi ini sebagai Goroutine (non-blocking)
	katakan("Halo")
}