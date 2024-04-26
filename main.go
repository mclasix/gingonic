package main

import (
	"gin/api"
	"net/http"
)

func main() {
	// Untuk menggunakan handler dari paket `api`, kamu dapat langsung memanggil fungsi Handler-nya.
	http.HandleFunc("/", api.Handler)

	// Mulai server HTTP
	http.ListenAndServe(":8080", nil)
}
