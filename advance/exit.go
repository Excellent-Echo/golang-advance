package advance

import (
	"fmt"
	"os"
)

func exit() {
	// exit
	// eksekusi code untuk menghentikan secara paksa
	// exit berada di dalam package "os"
	// didalam exit menerima argumen integer yang akan di tampilkan

	defer fmt.Println("finish")

	fmt.Println("start")
	fmt.Println("found some error")
	fmt.Println("fix error")
	os.Exit(1) // angkanya 0 , > 0 , 1 / 2 ... 100 exit status (int)
}
