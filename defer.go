package main

import "fmt"

func logging() {
	message := recover() // Menangkap error(panic), function ini seperti rescue dan catch
	fmt.Println("Error dengan message ", message)
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int, error bool) {
	defer logging() // Menjalankan function saat eksekusi selesai walaupun error (mirip finally dan ensure)

	if error {
		panic("Aplikasi Error") // Memberhentikan eksekusi program namun defer tetap menjalan (mirip raise dan throw)
	}

	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Hasil = ", result)
}

func main() {
	runApplication(0, true)
	fmt.Println("Semangat")
}
