package advance

import "fmt"

func orderSomeFood(food string) {
	defer fmt.Println("terima kasih, silahkan tunggu")

	fmt.Println("pesanan diterima")
	fmt.Println("pesanan anda :", food)
}

func advanceDefer() {
	// defer fmt.Println("finish")
	// defer
	// eksekusi code yang dijalankan sebelum eksekusi code selesai (di dalam func)

	// defer digabung dengan IIFE
	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		// IIFE
		func() {
			defer fmt.Println("halo 3")
		}()
	}

	fmt.Println("helo 2")

	// fmt.Println("start")

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("loop", i)
	// }
	// orderSomeFood("pizza")
	// orderSomeFood("burger")
}
