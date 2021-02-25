package main

import "fmt"

func endApp() {
	message := recover() // Menangkap error(panic), function ini seperti rescue dan catch
	fmt.Println("Error dengan message ", message)
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi Error") // Memberhentikan eksekusi program namun defer tetap menjalan (mirip raise dan throw)
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
}
